// Code generated by mockery v2.43.2. DO NOT EDIT.

package mockp2p

import (
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"

	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
)

// SubscriptionFilter is an autogenerated mock type for the SubscriptionFilter type
type SubscriptionFilter struct {
	mock.Mock
}

// CanSubscribe provides a mock function with given fields: _a0
func (_m *SubscriptionFilter) CanSubscribe(_a0 string) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CanSubscribe")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FilterIncomingSubscriptions provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionFilter) FilterIncomingSubscriptions(_a0 peer.ID, _a1 []*pubsub_pb.RPC_SubOpts) ([]*pubsub_pb.RPC_SubOpts, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FilterIncomingSubscriptions")
	}

	var r0 []*pubsub_pb.RPC_SubOpts
	var r1 error
	if rf, ok := ret.Get(0).(func(peer.ID, []*pubsub_pb.RPC_SubOpts) ([]*pubsub_pb.RPC_SubOpts, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(peer.ID, []*pubsub_pb.RPC_SubOpts) []*pubsub_pb.RPC_SubOpts); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pubsub_pb.RPC_SubOpts)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID, []*pubsub_pb.RPC_SubOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSubscriptionFilter creates a new instance of SubscriptionFilter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionFilter(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionFilter {
	mock := &SubscriptionFilter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
