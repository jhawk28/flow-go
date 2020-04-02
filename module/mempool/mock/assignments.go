// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import chunks "github.com/dapperlabs/flow-go/model/chunks"
import flow "github.com/dapperlabs/flow-go/model/flow"

import mock "github.com/stretchr/testify/mock"

// Assignments is an autogenerated mock type for the Assignments type
type Assignments struct {
	mock.Mock
}

// Add provides a mock function with given fields: assignmentFingerprint, assignment
func (_m *Assignments) Add(assignmentFingerprint flow.Identifier, assignment *chunks.Assignment) error {
	ret := _m.Called(assignmentFingerprint, assignment)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, *chunks.Assignment) error); ok {
		r0 = rf(assignmentFingerprint, assignment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// All provides a mock function with given fields:
func (_m *Assignments) All() []*chunks.Assignment {
	ret := _m.Called()

	var r0 []*chunks.Assignment
	if rf, ok := ret.Get(0).(func() []*chunks.Assignment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*chunks.Assignment)
		}
	}

	return r0
}

// ByID provides a mock function with given fields: assignmentID
func (_m *Assignments) ByID(assignmentID flow.Identifier) (*chunks.Assignment, error) {
	ret := _m.Called(assignmentID)

	var r0 *chunks.Assignment
	if rf, ok := ret.Get(0).(func(flow.Identifier) *chunks.Assignment); ok {
		r0 = rf(assignmentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chunks.Assignment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(assignmentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Has provides a mock function with given fields: assignmentID
func (_m *Assignments) Has(assignmentID flow.Identifier) bool {
	ret := _m.Called(assignmentID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(assignmentID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Rem provides a mock function with given fields: assignmentID
func (_m *Assignments) Rem(assignmentID flow.Identifier) bool {
	ret := _m.Called(assignmentID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(assignmentID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Size provides a mock function with given fields:
func (_m *Assignments) Size() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}
