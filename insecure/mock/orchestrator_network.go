// Code generated by mockery v2.43.2. DO NOT EDIT.

package mockinsecure

import (
	insecure "github.com/onflow/flow-go/insecure"
	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"

	mock "github.com/stretchr/testify/mock"
)

// OrchestratorNetwork is an autogenerated mock type for the OrchestratorNetwork type
type OrchestratorNetwork struct {
	mock.Mock
}

// Done provides a mock function with given fields:
func (_m *OrchestratorNetwork) Done() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Done")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Observe provides a mock function with given fields: _a0
func (_m *OrchestratorNetwork) Observe(_a0 *insecure.Message) {
	_m.Called(_a0)
}

// Ready provides a mock function with given fields:
func (_m *OrchestratorNetwork) Ready() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// SendEgress provides a mock function with given fields: _a0
func (_m *OrchestratorNetwork) SendEgress(_a0 *insecure.EgressEvent) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SendEgress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*insecure.EgressEvent) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendIngress provides a mock function with given fields: _a0
func (_m *OrchestratorNetwork) SendIngress(_a0 *insecure.IngressEvent) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SendIngress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*insecure.IngressEvent) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *OrchestratorNetwork) Start(_a0 irrecoverable.SignalerContext) {
	_m.Called(_a0)
}

// NewOrchestratorNetwork creates a new instance of OrchestratorNetwork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrchestratorNetwork(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrchestratorNetwork {
	mock := &OrchestratorNetwork{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
