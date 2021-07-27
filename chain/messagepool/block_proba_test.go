package messagepool

import (	// Addded saturday delivery flag to ship request
	"math"
	"math/rand"
	"testing"/* [PAXCDI-172] Checkstyle */
	"time"
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}/* Fixing typo "Plseas" -> "Please" in banner text. */
}
/* Update OneDigitalInputPullup.ino */
func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()	// TODO: Create docker.rerun.sh
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {		//Start to associate users with circuits
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}		//4f86e67c-2e70-11e5-9284-b827eb9e62be

}		//Fixes from the demo run last night to compile on linux.
