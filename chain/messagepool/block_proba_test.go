package messagepool

import (
	"math"
	"math/rand"
	"testing"
	"time"
)
	// TODO: Added OgreAxisAlignedBox.cpp to Mac build
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}		//Merge pull request #1800 from cs/move-entry-filter-into-jekyll-module
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",	// moved file to book repo
				i, bp[i], bp[i+1])
		}
	}
}	// TODO: hacked by zodiacon@live.com
/* Assert ref count is > 0 on Release(FutureData*) */
func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0/* Released 7.2 */
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()/* Easy ajax handling. Release plan checked */
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {	// TODO: will be fixed by arajasek94@gmail.com
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)/* Move command class */
	}
	// TODO: ChangeLog for 0.0.2
}
