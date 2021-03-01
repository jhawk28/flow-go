package trie

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common"
	"github.com/onflow/flow-go/ledger/common/utils"
	"github.com/onflow/flow-go/ledger/complete/mtrie/node"
)

// MTrie represents a perfect in-memory full binary Merkle tree with uniform height.
// For a detailed description of the storage model, please consult `mtrie/README.md`
//
// A MTrie is a thin wrapper around a the trie's root Node. An MTrie implements the
// logic for forming MTrie-graphs from the elementary nodes. Specifically:
//   * how Nodes (graph vertices) form a Trie,
//   * how register values are read from the trie,
//   * how Merkle proofs are generated from a trie, and
//   * how a new Trie with updated values is generated.
//
// `MTrie`s are _immutable_ data structures. Updating register values is implemented through
// copy-on-write, which creates a new `MTrie`. For minimal memory consumption, all sub-tries
// that where not affected by the write operation are shared between the original MTrie
// (before the register updates) and the updated MTrie (after the register writes).
//
// DEFINITIONS and CONVENTIONS:
//   * HEIGHT of a node v in a tree is the number of edges on the longest downward path
//     between v and a tree leaf. The height of a tree is the height of its root.
//     The height of a Trie is always the height of the fully-expanded tree.
type MTrie struct {
	root         *node.Node
	height       int
	pathByteSize int
}

// NewEmptyMTrie returns an empty Mtrie (root is an empty node)
func NewEmptyMTrie(pathByteSize int) (*MTrie, error) {
	if pathByteSize < 1 {
		return nil, errors.New("trie's path size [in bytes] must be positive")
	}
	height := pathByteSize * 8
	return &MTrie{
		root:         node.NewEmptyTreeRoot(height),
		pathByteSize: pathByteSize,
		height:       height,
	}, nil
}

// NewMTrie returns a Mtrie given the root
func NewMTrie(root *node.Node) (*MTrie, error) {
	if root.Height()%8 != 0 {
		return nil, errors.New("height of root node must be integer-multiple of 8")
	}
	pathByteSize := root.Height() / 8
	return &MTrie{
		root:         root,
		pathByteSize: pathByteSize,
		height:       root.Height(),
	}, nil
}

// StringRootHash returns the trie's Hex-encoded root hash.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) StringRootHash() string { return hex.EncodeToString(mt.root.Hash()) }

// RootHash returns the trie's root hash (i.e. the hash of the trie's root node).
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) RootHash() []byte { return mt.root.Hash() }

// PathLength return the length [in bytes] the trie operates with.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) PathLength() int { return mt.pathByteSize }

// AllocatedRegCount returns the number of allocated registers in the trie.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) AllocatedRegCount() uint64 { return mt.root.RegCount() }

// MaxDepth returns the length of the longest branch from root to leaf.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) MaxDepth() uint16 { return mt.root.MaxDepth() }

// RootNode returns the Trie's root Node
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) RootNode() *node.Node {
	return mt.root
}

// StringRootHash returns the trie's string representation.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) String() string {
	trieStr := fmt.Sprintf("Trie root hash: %v\n", mt.StringRootHash())
	return trieStr + mt.root.FmtStr("", "")
}

// UnsafeRead read payloads for the given paths. It is called unsafe as it requires the
// paths to be sorted
// TODO move consistency checks from Forrest into Trie to obtain a safe, self-contained API
func (mt *MTrie) UnsafeRead(paths []ledger.Path) []*ledger.Payload {
	res := make([]*ledger.Payload, 0, len(paths))
	mt.read(&res, mt.root, paths)
	return res
}

func (mt *MTrie) read(res *[]*ledger.Payload, head *node.Node, paths []ledger.Path) {
	// path not found
	if head == nil {
		for range paths {
			*res = append(*res, ledger.EmptyPayload())
		}
		return
	}
	// reached a leaf node
	if head.IsLeaf() {
		for _, p := range paths {
			if bytes.Equal(head.Path(), p) {
				*res = append(*res, head.Payload())
			} else {
				*res = append(*res, ledger.EmptyPayload())
			}
		}
		return
	}

	lpaths, rpaths := utils.SplitSortedPaths(paths, mt.height-head.Height())

	// must start with the left first, as left payloads have to be appended first
	if len(lpaths) > 0 {
		mt.read(res, head.LeftChild(), lpaths)
	}

	if len(rpaths) > 0 {
		mt.read(res, head.RightChild(), rpaths)
	}
}

// NewTrieWithUpdatedRegisters constructs a new trie containing all registers from the parent trie.
// The key-value pairs specify the registers whose values are supposed to hold updated values
// compared to the parent trie. Constructing the new trie is done in a COPY-ON-WRITE manner:
//   * The original trie remains unchanged.
//   * subtries that remain unchanged are from the parent trie instead of copied.
// UNSAFE: method requires the following conditions to be satisfied:
//   * keys are NOT duplicated
// TODO: move consistency checks from MForest to here, to make API safe and self-contained
func NewTrieWithUpdatedRegisters(parentTrie *MTrie, updatedPaths []ledger.Path, updatedPayloads []ledger.Payload) (*MTrie, error) {
	parentRoot := parentTrie.root
	updatedRoot := parentTrie.update(parentRoot.Height(), parentRoot, updatedPaths, updatedPayloads)
	updatedTrie, err := NewMTrie(updatedRoot)
	if err != nil {
		return nil, fmt.Errorf("constructing updated trie failed: %w", err)
	}
	return updatedTrie, nil
}

// update returns the head of updated sub-trie for the specified key-value pairs.
// UNSAFE: update requires the following conditions to be satisfied,
// but does not explicitly check them for performance reasons
//   * all keys AND the parent node share the same common prefix [0 : mt.maxHeight-1 - headHeight)
//     (excluding the bit at index headHeight)
//   * keys are NOT duplicated
func (parentTrie *MTrie) update(nodeHeight int, parentNode *node.Node, paths []ledger.Path, payloads []ledger.Payload) *node.Node {

	if len(paths) == 0 { // We are not changing any values in this sub-trie => return parent trie
		return parentNode
	}

	if parentNode == nil { // parent Trie has no sub-trie for the set of paths => construct entire subtree
		return parentTrie.constructSubtrie(nodeHeight, paths, payloads)
	}

	// from here on, we have parentNode != nil AND len(paths) > 0
	if parentNode.IsLeaf() { // parent node is a leaf, i.e. parent Trie only stores a single value in this sub-trie
		parentPath := parentNode.Path() // Per definition, a leaf must have a non-nil path
		for _, p := range paths {       // TODO: binary search if paths are sorted - or maybe inputs paths->payload should be a map?
			if bytes.Equal(p, parentPath) {
				return parentTrie.constructSubtrie(nodeHeight, paths, payloads)
			}
		}
		// TODO: copy payload when using in-place MergeSort for separating the payloads
		paths = append(paths, parentNode.Path()) // TODO: if paths are sorted, insert in order
		payloads = append(payloads, *parentNode.Payload())

		return parentTrie.constructSubtrie(nodeHeight, paths, payloads)
	}

	// Split payloads so we can update the trie in parallel
	lpaths, lpayloads, rpaths, rpayloads := utils.SplitByPath(paths, payloads, parentTrie.height-nodeHeight)

	// TODO [runtime optimization]: do not branch if either lpayload or rpayload is empty
	// TODO: what does the above TODO mean?
	var lChild, rChild *node.Node
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		lChild = parentTrie.update(nodeHeight-1, parentNode.LeftChild(), lpaths, lpayloads)
	}()
	rChild = parentTrie.update(nodeHeight-1, parentNode.RightChild(), rpaths, rpayloads)
	wg.Wait()

	if lChild == parentNode.LeftChild() && rChild == parentNode.RightChild() {
		return parentNode
	}
	return node.NewInterimNode(nodeHeight, lChild, rChild)
}

// constructSubtrie returns the head of a newly-constructed sub-trie for the specified key-value pairs.
// UNSAFE: constructSubtrie requires the following conditions to be satisfied,
// but does not explicitly check them for performance reasons
//   * paths all share the same common prefix [0 : mt.maxHeight-1 - headHeight)
//     (excluding the bit at index headHeight)
//   * paths contains at least one element
//   * paths are NOT duplicated
// TODO: remove error return
func (parentTrie *MTrie) constructSubtrie(nodeHeight int, paths []ledger.Path, payloads []ledger.Payload) *node.Node {
	// no inserts => default value, represented by nil node
	if len(paths) == 0 {
		return nil
	}
	// If we are at a leaf node, we create the node
	if len(paths) == 1 {
		return node.NewLeaf(paths[0], &payloads[0], nodeHeight)
	}
	// from here on, we have: len(paths) > 1

	// Split updates by paths so we can update the trie in parallel
	lpaths, lpayloads, rpaths, rpayloads := utils.SplitByPath(paths, payloads, parentTrie.height-nodeHeight)
	// Note: (pathLength-height) will never reach the value pathLength, i.e. we will never execute this code for height==0
	// This is because at height=0, we only have (at most) one path left, as paths are not duplicated
	// (by requirement of this function). But even if this condition is violated, the code will not return a faulty
	// but instead panic with Index Out Of Range error

	// TODO [runtime optimization]: do not branch if either lpaths or rpaths is empty
	var lChild, rChild *node.Node
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		lChild = parentTrie.constructSubtrie(nodeHeight-1, lpaths, lpayloads)
	}()
	rChild = parentTrie.constructSubtrie(nodeHeight-1, rpaths, rpayloads)
	wg.Wait()

	return node.NewInterimNode(nodeHeight, lChild, rChild)
}

// UnsafeProofs provides proofs for the given paths, this is called unsafe as
// it requires the input paths to be sorted in advance.
func (mt *MTrie) UnsafeProofs(paths []ledger.Path, proofs []*ledger.TrieProof) {
	mt.proofs(mt.root, paths, proofs)
}

func (mt *MTrie) proofs(head *node.Node, paths []ledger.Path, proofs []*ledger.TrieProof) {
	// we've reached the end of a trie
	// and path is not found (noninclusion proof)
	if head == nil {
		return
	}

	// we've reached a leaf
	if head.IsLeaf() {
		// value matches (inclusion proof)
		if bytes.Equal(head.Path(), paths[0]) {
			proofs[0].Path = head.Path()
			proofs[0].Payload = head.Payload()
			proofs[0].Inclusion = true
		}
		// TODO: insert ERROR if len(paths) != 1
		return
	}

	// increment steps for all the proofs
	for _, p := range proofs {
		p.Steps++
	}
	// split paths based on the value of i-th bit (i = trie height - node height)
	lpaths, lproofs, rpaths, rproofs := utils.SplitTrieProofsByPath(paths, proofs, mt.height-head.Height())

	if len(lpaths) > 0 {
		if rChild := head.RightChild(); rChild != nil { // TODO: is that a sanity check?
			nodeHash := rChild.Hash()
			isDef := bytes.Equal(nodeHash, common.GetDefaultHashForHeight(rChild.Height())) // TODO: why not rChild.RegisterCount != 0?
			if !isDef {                                                                     // in proofs, we only provide non-default value hashes
				for _, p := range lproofs {
					utils.SetBit(p.Flags, mt.height-head.Height())
					p.Interims = append(p.Interims, nodeHash)
				}
			}
		}
		mt.proofs(head.LeftChild(), lpaths, lproofs)
	}

	if len(rpaths) > 0 {
		if lChild := head.LeftChild(); lChild != nil {
			nodeHash := lChild.Hash()
			isDef := bytes.Equal(nodeHash, common.GetDefaultHashForHeight(lChild.Height()))
			if !isDef { // in proofs, we only provide non-default value hashes
				for _, p := range rproofs {
					utils.SetBit(p.Flags, mt.height-head.Height())
					p.Interims = append(p.Interims, nodeHash)
				}
			}
		}
		mt.proofs(head.RightChild(), rpaths, rproofs)
	}
}

// Equals compares two tries for equality.
// Tries are equal iff they store the same data (i.e. root hash matches)
// and their number and height are identical
func (mt *MTrie) Equals(o *MTrie) bool {
	if o == nil {
		return false
	}
	return o.PathLength() == mt.PathLength() && bytes.Equal(o.RootHash(), mt.RootHash())
}

// DumpAsJSON dumps the trie key value pairs to a file having each key value pair as a json row
func (mt *MTrie) DumpAsJSON(w io.Writer) error {

	// Use encoder to prevent building entire trie in memory
	enc := json.NewEncoder(w)

	err := mt.dumpAsJSON(mt.root, enc)
	if err != nil {
		return err
	}

	return nil
}

func (mt *MTrie) dumpAsJSON(n *node.Node, encoder *json.Encoder) error {
	if n.IsLeaf() {
		err := encoder.Encode(n.Payload())
		if err != nil {
			return err
		}
		return nil
	}

	if lChild := n.LeftChild(); lChild != nil {
		err := mt.dumpAsJSON(lChild, encoder)
		if err != nil {
			return err
		}
	}

	if rChild := n.RightChild(); rChild != nil {
		err := mt.dumpAsJSON(rChild, encoder)
		if err != nil {
			return err
		}
	}
	return nil
}

// EmptyTrieRootHash returns the rootHash of an empty Trie for the specified path size [bytes]
func EmptyTrieRootHash(pathByteSize int) []byte {
	return node.NewEmptyTreeRoot(8 * pathByteSize).Hash()
}

// AllPayloads returns all payloads
func (mt *MTrie) AllPayloads() []ledger.Payload {
	return mt.root.AllPayloads()
}

// IsAValidTrie verifies the content of the trie for potential issues
func (mt *MTrie) IsAValidTrie() bool {
	// TODO add checks on the health of node max height ...
	return mt.root.VerifyCachedHash()
}
