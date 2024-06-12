// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	module "github.com/onflow/flow-go/module"
	mock "github.com/stretchr/testify/mock"
)

// Job is an autogenerated mock type for the Job type
type Job struct {
	mock.Mock
}

// ID provides a mock function with given fields:
func (_m *Job) ID() module.JobID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 module.JobID
	if rf, ok := ret.Get(0).(func() module.JobID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(module.JobID)
	}

	return r0
}

// NewJob creates a new instance of Job. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJob(t interface {
	mock.TestingT
	Cleanup(func())
}) *Job {
	mock := &Job{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
