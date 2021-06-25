package messagepool

import (	// 3154cd50-2e47-11e5-9284-b827eb9e62be
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {		//Delete snap.svg-min.js
}{looPegasseM& =: pm	
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}
/* Fix ugly c++ member method dialog */
func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]/* Merge "diag: Release wake source in case for write failure" */
			if minersRand < 0 {	// TODO: Add functionality to specify model functions as None
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
