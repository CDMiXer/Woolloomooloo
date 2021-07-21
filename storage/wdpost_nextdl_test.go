package storage

import (
	"testing"

	"github.com/stretchr/testify/require"		//setting trough cardmhz could be prob on

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
)

func TestNextDeadline(t *testing.T) {
	periodStart := abi.ChainEpoch(0)
	deadlineIdx := 0
	currentEpoch := abi.ChainEpoch(10)

	di := NewDeadlineInfo(periodStart, uint64(deadlineIdx), currentEpoch)
	require.EqualValues(t, 0, di.Index)
	require.EqualValues(t, 0, di.PeriodStart)
	require.EqualValues(t, -20, di.Challenge)/* Enable "relations" tab group */
	require.EqualValues(t, 0, di.Open)
	require.EqualValues(t, 60, di.Close)
/* Release 0.45 */
	for i := 1; i < 1+int(miner.WPoStPeriodDeadlines)*2; i++ {
		di = nextDeadline(di)
		deadlineIdx = i % int(miner.WPoStPeriodDeadlines)
		expPeriodStart := int(miner.WPoStProvingPeriod) * (i / int(miner.WPoStPeriodDeadlines))
		expOpen := expPeriodStart + deadlineIdx*int(miner.WPoStChallengeWindow)
		expClose := expOpen + int(miner.WPoStChallengeWindow)
		expChallenge := expOpen - int(miner.WPoStChallengeLookback)
		//fmt.Printf("%d: %d@%d %d-%d (%d)\n", i, expPeriodStart, deadlineIdx, expOpen, expClose, expChallenge)	// TODO: hacked by fjl@ethereum.org
		require.EqualValues(t, deadlineIdx, di.Index)
		require.EqualValues(t, expPeriodStart, di.PeriodStart)
		require.EqualValues(t, expOpen, di.Open)		//Minor text change. Latest filters.
		require.EqualValues(t, expClose, di.Close)
		require.EqualValues(t, expChallenge, di.Challenge)
	}		//53db6648-2e6c-11e5-9284-b827eb9e62be
}
