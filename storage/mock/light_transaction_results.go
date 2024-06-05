// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/onflow/flow-go/storage"
)

// LightTransactionResults is an autogenerated mock type for the LightTransactionResults type
type LightTransactionResults struct {
	mock.Mock
}

// BatchStore provides a mock function with given fields: blockID, transactionResults, batch
func (_m *LightTransactionResults) BatchStore(blockID flow.Identifier, transactionResults []flow.LightTransactionResult, batch storage.BatchStorage) error {
	ret := _m.Called(blockID, transactionResults, batch)

	if len(ret) == 0 {
		panic("no return value specified for BatchStore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, []flow.LightTransactionResult, storage.BatchStorage) error); ok {
		r0 = rf(blockID, transactionResults, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ByBlockID provides a mock function with given fields: id
func (_m *LightTransactionResults) ByBlockID(id flow.Identifier) ([]flow.LightTransactionResult, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ByBlockID")
	}

	var r0 []flow.LightTransactionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) ([]flow.LightTransactionResult, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) []flow.LightTransactionResult); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.LightTransactionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ByBlockIDTransactionID provides a mock function with given fields: blockID, transactionID
func (_m *LightTransactionResults) ByBlockIDTransactionID(blockID flow.Identifier, transactionID flow.Identifier) (*flow.LightTransactionResult, error) {
	ret := _m.Called(blockID, transactionID)

	if len(ret) == 0 {
		panic("no return value specified for ByBlockIDTransactionID")
	}

	var r0 *flow.LightTransactionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.Identifier) (*flow.LightTransactionResult, error)); ok {
		return rf(blockID, transactionID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.Identifier) *flow.LightTransactionResult); ok {
		r0 = rf(blockID, transactionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.LightTransactionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier, flow.Identifier) error); ok {
		r1 = rf(blockID, transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ByBlockIDTransactionIndex provides a mock function with given fields: blockID, txIndex
func (_m *LightTransactionResults) ByBlockIDTransactionIndex(blockID flow.Identifier, txIndex uint32) (*flow.LightTransactionResult, error) {
	ret := _m.Called(blockID, txIndex)

	if len(ret) == 0 {
		panic("no return value specified for ByBlockIDTransactionIndex")
	}

	var r0 *flow.LightTransactionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, uint32) (*flow.LightTransactionResult, error)); ok {
		return rf(blockID, txIndex)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier, uint32) *flow.LightTransactionResult); ok {
		r0 = rf(blockID, txIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.LightTransactionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier, uint32) error); ok {
		r1 = rf(blockID, txIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLightTransactionResults creates a new instance of LightTransactionResults. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLightTransactionResults(t interface {
	mock.TestingT
	Cleanup(func())
}) *LightTransactionResults {
	mock := &LightTransactionResults{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
