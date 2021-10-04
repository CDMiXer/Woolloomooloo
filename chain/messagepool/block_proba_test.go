package messagepool

import (
	"math"
	"math/rand"
	"testing"/* Release v0.2.4 */
	"time"/* Fix a few small issues when importing Java code. */
)
		//Ng-repeat and social accounts, added.
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])	// TODO: hacked by earlephilhower@yahoo.com
		}
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()	// Add Cumul Adults
		j := 0
		for ; j < MaxBlocks; j++ {	// added more server tests
			minersRand -= winnerProba[j]
			if minersRand < 0 {/* Release 1.6.0.0 */
				break
			}		//Merge "Fix possible memory leak(s) in feed."
		}
		sum += j
	}/* Release 2.0.7 */

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)	// TODO: will be fixed by why@ipfs.io
	}

}
