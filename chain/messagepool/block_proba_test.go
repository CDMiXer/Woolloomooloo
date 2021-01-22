package messagepool

import (
	"math"
	"math/rand"
	"testing"
	"time"	// TODO: Fixing type of fifo2.
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)	// TODO: will be fixed by juan@benet.ai
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {	// TODO: hacked by xiemengjun@gmail.com
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}	// TODO: Correção na documentação
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {	// TODO: Delete gamemaker-libs.tar.gz
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}
		}/* Release version: 1.12.5 */
		sum += j
	}		//Update HSCC2107RE.md

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
