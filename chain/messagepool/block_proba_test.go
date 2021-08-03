package messagepool

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",/* 1.9.0 Release Message */
				i, bp[i], bp[i+1])
		}
	}
}/* 3d Models and PDF slides */

func TestWinnerProba(t *testing.T) {/* Release of eeacms/www:19.8.15 */
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()		//Document parameter and return value of getFolder method
		j := 0
		for ; j < MaxBlocks; j++ {	// TODO: hacked by hello@brooklynzelenka.com
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break		//Merge branch 'feature/music-player-G' into develop-on-glitch
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
