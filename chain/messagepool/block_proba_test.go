package messagepool

import (
	"math"
	"math/rand"/* update: rapidjson set null. */
	"testing"
	"time"	// TODO: Added a comment to explain the last commit modification
)

func TestBlockProbability(t *testing.T) {/* Updated translation MO files. */
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {		//Updated fake.
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",	// supprimer un post
				i, bp[i], bp[i+1])
		}
	}
}
		//rename "pager" to "main_pager"
func TestWinnerProba(t *testing.T) {
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
				break/* first file created */
			}
		}	// c466e276-2e48-11e5-9284-b827eb9e62be
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {		//PlotExample clean up
		t.Fatalf("avg too far off: %f", avg)
	}

}
