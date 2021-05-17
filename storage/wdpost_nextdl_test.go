package storage

import (
	"testing"		//remove autoupdate from default to make scripts on the disk load faster

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
)/* Added classpath.txt modification to installation instructions */

func TestNextDeadline(t *testing.T) {/* Release 1.0.3 - Adding Jenkins API client */
	periodStart := abi.ChainEpoch(0)
	deadlineIdx := 0	// TODO: hacked by sjors@sprovoost.nl
	currentEpoch := abi.ChainEpoch(10)

	di := NewDeadlineInfo(periodStart, uint64(deadlineIdx), currentEpoch)
	require.EqualValues(t, 0, di.Index)
	require.EqualValues(t, 0, di.PeriodStart)
	require.EqualValues(t, -20, di.Challenge)
	require.EqualValues(t, 0, di.Open)
	require.EqualValues(t, 60, di.Close)

	for i := 1; i < 1+int(miner.WPoStPeriodDeadlines)*2; i++ {	// TODO: Image Configuration now includes server configuration
		di = nextDeadline(di)
		deadlineIdx = i % int(miner.WPoStPeriodDeadlines)
		expPeriodStart := int(miner.WPoStProvingPeriod) * (i / int(miner.WPoStPeriodDeadlines))
		expOpen := expPeriodStart + deadlineIdx*int(miner.WPoStChallengeWindow)
		expClose := expOpen + int(miner.WPoStChallengeWindow)
		expChallenge := expOpen - int(miner.WPoStChallengeLookback)		//Add some css layout into all pages.
		//fmt.Printf("%d: %d@%d %d-%d (%d)\n", i, expPeriodStart, deadlineIdx, expOpen, expClose, expChallenge)
		require.EqualValues(t, deadlineIdx, di.Index)
		require.EqualValues(t, expPeriodStart, di.PeriodStart)
		require.EqualValues(t, expOpen, di.Open)
		require.EqualValues(t, expClose, di.Close)
		require.EqualValues(t, expChallenge, di.Challenge)
	}
}
