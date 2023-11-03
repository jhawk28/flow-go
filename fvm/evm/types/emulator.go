package types

import (
	"math/big"

	gethCommon "github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	DefaultDirectCallBaseGasUsage = uint64(21_000)
	DefaultDirectCallGasPrice     = uint64(0)

	// anything block number above 0 works here
	BlockNumberForEVMRules = big.NewInt(1)
)

// BlockContext holds the context needed for the emulator operations
type BlockContext struct {
	BlockNumber            uint64
	DirectCallBaseGasUsage uint64
	DirectCallGasPrice     uint64
	GasFeeCollector        Address
}

// NewDefaultBlockContext returns a new default block context
func NewDefaultBlockContext(BlockNumber uint64) BlockContext {
	return BlockContext{
		BlockNumber:            BlockNumber,
		DirectCallBaseGasUsage: DefaultDirectCallBaseGasUsage,
		DirectCallGasPrice:     DefaultDirectCallGasPrice,
	}
}

// ReadOnlyBlockView provides a read only view of a block
type ReadOnlyBlockView interface {
	// BalanceOf returns the balance of this address
	BalanceOf(address Address) (*big.Int, error)
	// NonceOf returns the nonce of this address
	NonceOf(address Address) (uint64, error)
	// CodeOf returns the code for this address (if smart contract is deployed at this address)
	CodeOf(address Address) (Code, error)
}

// BlockView allows evm calls in the context of a block
// Errors returned by BlockView is one of these types
// - Fatal error
// - Database error (non-fatal)
// - EVM validation error
// - EVM execution error
type BlockView interface {
	// executes a direct call
	DirectCall(call *DirectCall) (*Result, error)

	// RunTransaction executes an evm transaction
	RunTransaction(tx *gethTypes.Transaction) (*Result, error)
}

type Emulator interface {
	// constructs a new block view
	NewReadOnlyBlockView(ctx BlockContext) (ReadOnlyBlockView, error)

	// constructs a new block
	NewBlockView(ctx BlockContext) (BlockView, error)
}

// Errors returned by Database is one of these types
// - Fatal error
// - Database error (non-fatal)
type Database interface {
	ethdb.KeyValueStore

	// Commit commits the changes
	Commit(rootHash gethCommon.Hash) error

	// GetRootHash returns the active root hash
	GetRootHash() (gethCommon.Hash, error)
}
