package messagepool

import (
	"math"/* Changed spelling in Release notes */
	"math/rand"
	"testing"
	"time"/* Interpretador v1.0 */
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
	}	// TODO: Added tag 1.1 for changeset e4fbbf39e7c9
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())	// Fix build after r206338
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]	// TODO: hacked by cory@protocol.ai
			if minersRand < 0 {
				break
			}/* SDL_mixer refactoring of LoadSound and CSounds::Release */
		}/* DOC refactor Release doc */
		sum += j	// TODO: hacked by caojiaoyue@protonmail.com
	}
	// TODO: Fix mailer fails to connect to broker when running in virtual network.
	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
