package cohort2

import (
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/utils/unittest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/onflow/flow-go/integration/tests/epochs"
)

func TestEpochJoinAndLeaveVN(t *testing.T) {
	suite.Run(t, new(EpochJoinAndLeaveVNSuite))
}

type EpochJoinAndLeaveVNSuite struct {
	epochs.DynamicEpochTransitionSuite
}

func (s *EpochJoinAndLeaveVNSuite) SetupTest() {
	// require approvals for seals to verify that the joining VN is producing valid seals in the second epoch
	s.RequiredSealApprovals = 1
	// increase epoch length to account for greater sealing lag due to above
	// NOTE: this value is set fairly aggressively to ensure shorter test times.
	// If flakiness due to failure to complete staking operations in time is observed,
	// try increasing (by 10-20 views).
	s.StakingAuctionLen = 100
	s.DKGPhaseLen = 100
	s.EpochLen = 450
	s.EpochCommitSafetyThreshold = 20
	s.DynamicEpochTransitionSuite.SetupTest()
}

// TestEpochJoinAndLeaveVN should update verification nodes and assert healthy network conditions
// after the epoch transition completes. See health check function for details.
func (s *EpochJoinAndLeaveVNSuite) TestEpochJoinAndLeaveVN() {
	unittest.SkipUnless(s.T(), unittest.TEST_TODO, "flaky")
	s.RunTestEpochJoinAndLeave(flow.RoleVerification, s.AssertNetworkHealthyAfterVNChange)
}
