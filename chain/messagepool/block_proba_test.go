package messagepool

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {		//Added support for new constructor of ProxyConfiguration
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {	// TODO: :memo: APP #129
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}/* fix errors spotted by pychecker */
	}
}/* whoa fix that scrollbar halving */
/* Merge "wlan: Release 3.2.3.108" */
func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000/* Moved core sources and pom into top level project */
	winnerProba := noWinnersProb()		//Changed url back to http://ikangiec.github.io
	sum := 0/* Merge "Release notes for implied roles" */
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()/* Minor grammar fix at the start of the README */
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break	// TODO: Merge "FIX for compute scale down"
			}
		}
		sum += j
	}	// added simple helper default classes

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
