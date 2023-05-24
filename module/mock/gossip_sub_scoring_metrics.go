// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	channels "github.com/onflow/flow-go/network/channels"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// GossipSubScoringMetrics is an autogenerated mock type for the GossipSubScoringMetrics type
type GossipSubScoringMetrics struct {
	mock.Mock
}

// OnAppSpecificScoreUpdated provides a mock function with given fields: _a0
func (_m *GossipSubScoringMetrics) OnAppSpecificScoreUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnBehaviourPenaltyUpdated provides a mock function with given fields: _a0
func (_m *GossipSubScoringMetrics) OnBehaviourPenaltyUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnFirstMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *GossipSubScoringMetrics) OnFirstMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnIPColocationFactorUpdated provides a mock function with given fields: _a0
func (_m *GossipSubScoringMetrics) OnIPColocationFactorUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnInvalidMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *GossipSubScoringMetrics) OnInvalidMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnMeshMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *GossipSubScoringMetrics) OnMeshMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnOverallPeerScoreUpdated provides a mock function with given fields: _a0
func (_m *GossipSubScoringMetrics) OnOverallPeerScoreUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnTimeInMeshUpdated provides a mock function with given fields: _a0, _a1
func (_m *GossipSubScoringMetrics) OnTimeInMeshUpdated(_a0 channels.Topic, _a1 time.Duration) {
	_m.Called(_a0, _a1)
}

// SetWarningStateCount provides a mock function with given fields: _a0
func (_m *GossipSubScoringMetrics) SetWarningStateCount(_a0 uint) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewGossipSubScoringMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewGossipSubScoringMetrics creates a new instance of GossipSubScoringMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGossipSubScoringMetrics(t mockConstructorTestingTNewGossipSubScoringMetrics) *GossipSubScoringMetrics {
	mock := &GossipSubScoringMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
