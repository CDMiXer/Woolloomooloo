package messagepool

import (/* Release preparation */
	"math"
	"math/rand"
	"testing"
	"time"/* Release Django Evolution 0.6.4. */
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}	// TODO: hacked by boringland@protonmail.ch
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}
/* fix tests print */
func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())	// TODO: début bulletin json
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()/* Deleted msmeter2.0.1/Release/meter.obj */
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {	// TODO: will be fixed by martin2cai@hotmail.com
		t.Fatalf("avg too far off: %f", avg)
	}

}
