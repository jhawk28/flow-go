// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	context "context"

	derived "github.com/onflow/flow-go/fvm/derived"
	entity "github.com/onflow/flow-go/module/mempool/entity"

	execution "github.com/onflow/flow-go/engine/execution"

	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	state "github.com/onflow/flow-go/fvm/state"
)

// BlockComputer is an autogenerated mock type for the BlockComputer type
type BlockComputer struct {
	mock.Mock
}

// ExecuteBlock provides a mock function with given fields: ctx, parentBlockExecutionResultID, block, view, derivedBlockData
func (_m *BlockComputer) ExecuteBlock(ctx context.Context, parentBlockExecutionResultID flow.Identifier, block *entity.ExecutableBlock, view state.View, derivedBlockData *derived.DerivedBlockData) (*execution.ComputationResult, error) {
	ret := _m.Called(ctx, parentBlockExecutionResultID, block, view, derivedBlockData)

	var r0 *execution.ComputationResult
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier, *entity.ExecutableBlock, state.View, *derived.DerivedBlockData) *execution.ComputationResult); ok {
		r0 = rf(ctx, parentBlockExecutionResultID, block, view, derivedBlockData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.ComputationResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier, *entity.ExecutableBlock, state.View, *derived.DerivedBlockData) error); ok {
		r1 = rf(ctx, parentBlockExecutionResultID, block, view, derivedBlockData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBlockComputer interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlockComputer creates a new instance of BlockComputer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlockComputer(t mockConstructorTestingTNewBlockComputer) *BlockComputer {
	mock := &BlockComputer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
