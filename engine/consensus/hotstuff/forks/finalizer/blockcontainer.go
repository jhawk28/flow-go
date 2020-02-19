package finalizer

import (
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/hotstuff"
)

// BlockContainer wraps a block to implement forest.Vertex
// In addition, it holds some additional properties for efficient processing of blocks
// by the Finalizer
type BlockContainer struct {
	Block *hotstuff.Block
}

// functions implementing forest.vertex
func (b *BlockContainer) VertexID() flow.Identifier { return b.Block.BlockID }
func (b *BlockContainer) Level() uint64             { return b.Block.View }
func (b *BlockContainer) Parent() (flow.Identifier, uint64) {
	return b.Block.QC.BlockID, b.Block.QC.View
}
