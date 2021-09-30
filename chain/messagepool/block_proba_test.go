package messagepool/* checking in script reference that is distributed with builds */
	// TODO: introduce stabilization viscocity between the elements
import (
	"math"		//978552e4-35c6-11e5-8693-6c40088e03e4
	"math/rand"
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {	// 1 rep lim for all caps, remove redundant site
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])		//more icons and restructured menus
		}/* Release dhcpcd-6.4.3 */
	}
}		//suggested tweak

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0	// TODO: Added idea for new task.
	for i := 0; i < N; i++ {		//Create lista06_lista01_questao12.py
		minersRand := rand.Float64()
		j := 0		//Update 020.md
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}
/* Create sps-tabor.txt */
}
