// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"

	time "time"
)

// CommunicatorConsumer is an autogenerated mock type for the CommunicatorConsumer type
type CommunicatorConsumer struct {
	mock.Mock
}

// OnOwnProposal provides a mock function with given fields: proposal, targetPublicationTime
func (_m *CommunicatorConsumer) OnOwnProposal(proposal *flow.Header, targetPublicationTime time.Time) {
	_m.Called(proposal, targetPublicationTime)
}

// OnOwnTimeout provides a mock function with given fields: timeout
func (_m *CommunicatorConsumer) OnOwnTimeout(timeout *model.TimeoutObject) {
	_m.Called(timeout)
}

// OnOwnVote provides a mock function with given fields: blockID, view, sigData, recipientID
func (_m *CommunicatorConsumer) OnOwnVote(blockID flow.Identifier, view uint64, sigData []byte, recipientID flow.Identifier) {
	_m.Called(blockID, view, sigData, recipientID)
}

type mockConstructorTestingTNewCommunicatorConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommunicatorConsumer creates a new instance of CommunicatorConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommunicatorConsumer(t mockConstructorTestingTNewCommunicatorConsumer) *CommunicatorConsumer {
	mock := &CommunicatorConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
