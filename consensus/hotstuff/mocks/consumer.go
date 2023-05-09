// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	hotstuff "github.com/onflow/flow-go/consensus/hotstuff"
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"

	time "time"
)

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// OnBlockIncorporated provides a mock function with given fields: _a0
func (_m *Consumer) OnBlockIncorporated(_a0 *model.Block) {
	_m.Called(_a0)
}

// OnCurrentViewDetails provides a mock function with given fields: currentView, finalizedView, currentLeader
func (_m *Consumer) OnCurrentViewDetails(currentView uint64, finalizedView uint64, currentLeader flow.Identifier) {
	_m.Called(currentView, finalizedView, currentLeader)
}

// OnDoubleProposeDetected provides a mock function with given fields: _a0, _a1
func (_m *Consumer) OnDoubleProposeDetected(_a0 *model.Block, _a1 *model.Block) {
	_m.Called(_a0, _a1)
}

// OnEventProcessed provides a mock function with given fields:
func (_m *Consumer) OnEventProcessed() {
	_m.Called()
}

// OnFinalizedBlock provides a mock function with given fields: _a0
func (_m *Consumer) OnFinalizedBlock(_a0 *model.Block) {
	_m.Called(_a0)
}

// OnInvalidBlockDetected provides a mock function with given fields: err
func (_m *Consumer) OnInvalidBlockDetected(err model.InvalidBlockError) {
	_m.Called(err)
}

// OnLocalTimeout provides a mock function with given fields: currentView
func (_m *Consumer) OnLocalTimeout(currentView uint64) {
	_m.Called(currentView)
}

// OnOwnProposal provides a mock function with given fields: proposal, targetPublicationTime
func (_m *Consumer) OnOwnProposal(proposal *flow.Header, targetPublicationTime time.Time) {
	_m.Called(proposal, targetPublicationTime)
}

// OnOwnTimeout provides a mock function with given fields: timeout
func (_m *Consumer) OnOwnTimeout(timeout *model.TimeoutObject) {
	_m.Called(timeout)
}

// OnOwnVote provides a mock function with given fields: blockID, view, sigData, recipientID
func (_m *Consumer) OnOwnVote(blockID flow.Identifier, view uint64, sigData []byte, recipientID flow.Identifier) {
	_m.Called(blockID, view, sigData, recipientID)
}

// OnPartialTc provides a mock function with given fields: currentView, partialTc
func (_m *Consumer) OnPartialTc(currentView uint64, partialTc *hotstuff.PartialTcCreated) {
	_m.Called(currentView, partialTc)
}

// OnQcTriggeredViewChange provides a mock function with given fields: oldView, newView, qc
func (_m *Consumer) OnQcTriggeredViewChange(oldView uint64, newView uint64, qc *flow.QuorumCertificate) {
	_m.Called(oldView, newView, qc)
}

// OnReceiveProposal provides a mock function with given fields: currentView, proposal
func (_m *Consumer) OnReceiveProposal(currentView uint64, proposal *model.Proposal) {
	_m.Called(currentView, proposal)
}

// OnReceiveQc provides a mock function with given fields: currentView, qc
func (_m *Consumer) OnReceiveQc(currentView uint64, qc *flow.QuorumCertificate) {
	_m.Called(currentView, qc)
}

// OnReceiveTc provides a mock function with given fields: currentView, tc
func (_m *Consumer) OnReceiveTc(currentView uint64, tc *flow.TimeoutCertificate) {
	_m.Called(currentView, tc)
}

// OnStart provides a mock function with given fields: currentView
func (_m *Consumer) OnStart(currentView uint64) {
	_m.Called(currentView)
}

// OnStartingTimeout provides a mock function with given fields: _a0
func (_m *Consumer) OnStartingTimeout(_a0 model.TimerInfo) {
	_m.Called(_a0)
}

// OnTcTriggeredViewChange provides a mock function with given fields: oldView, newView, tc
func (_m *Consumer) OnTcTriggeredViewChange(oldView uint64, newView uint64, tc *flow.TimeoutCertificate) {
	_m.Called(oldView, newView, tc)
}

// OnViewChange provides a mock function with given fields: oldView, newView
func (_m *Consumer) OnViewChange(oldView uint64, newView uint64) {
	_m.Called(oldView, newView)
}

type mockConstructorTestingTNewConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumer creates a new instance of Consumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumer(t mockConstructorTestingTNewConsumer) *Consumer {
	mock := &Consumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
