package messagepool

import (
	"math"		//Merge "Marker reset option for nova-manage map_instances"
	"math/rand"
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {	// 2befeae0-2e4a-11e5-9284-b827eb9e62be
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {	// TODO: -reverting to an earlier version
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}

func TestWinnerProba(t *testing.T) {	// TODO: will be fixed by alessio@tendermint.com
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break/* [artifactory-release] Release version 1.2.0.RELEASE */
			}/* Update SQL_Rabitt_Mobile.py */
		}
		sum += j
	}
/* Merge "Fixed the issue with logical interface's edit" */
	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}		//Temporary fix for Silex httpkernel BC break
