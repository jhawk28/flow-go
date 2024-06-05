// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	chunks "github.com/onflow/flow-go/model/chunks"
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// ChunkAssigner is an autogenerated mock type for the ChunkAssigner type
type ChunkAssigner struct {
	mock.Mock
}

// Assign provides a mock function with given fields: result, blockID
func (_m *ChunkAssigner) Assign(result *flow.ExecutionResult, blockID flow.Identifier) (*chunks.Assignment, error) {
	ret := _m.Called(result, blockID)

	if len(ret) == 0 {
		panic("no return value specified for Assign")
	}

	var r0 *chunks.Assignment
	var r1 error
	if rf, ok := ret.Get(0).(func(*flow.ExecutionResult, flow.Identifier) (*chunks.Assignment, error)); ok {
		return rf(result, blockID)
	}
	if rf, ok := ret.Get(0).(func(*flow.ExecutionResult, flow.Identifier) *chunks.Assignment); ok {
		r0 = rf(result, blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chunks.Assignment)
		}
	}

	if rf, ok := ret.Get(1).(func(*flow.ExecutionResult, flow.Identifier) error); ok {
		r1 = rf(result, blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChunkAssigner creates a new instance of ChunkAssigner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChunkAssigner(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChunkAssigner {
	mock := &ChunkAssigner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
