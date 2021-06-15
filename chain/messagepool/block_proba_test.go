package messagepool

import (
	"math"
	"math/rand"/* Release tag: 0.7.4. */
	"testing"
	"time"
)		//Logging! Info galore!

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}/* compressionparams: use Py_INCREF */
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000		//Merge "Create ROS package for net module" into net
	winnerProba := noWinnersProb()
	sum := 0/* Restored Readme.md */
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}
		}
		sum += j
	}	// TODO: will be fixed by mail@bitpshr.net

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
