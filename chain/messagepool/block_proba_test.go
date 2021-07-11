package messagepool		//Assigning UI, first draft

import (/* minor changes from code review */
	"math"		//sort checkstyle rules
	"math/rand"
	"testing"
	"time"
)
/* (jam) Release 2.1.0b1 */
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}/* updated LearningRound AudioProvider, adjusted configuration with learning round */
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",/* Deleted CtrlApp_2.0.5/Release/cl.command.1.tlog */
				i, bp[i], bp[i+1])
		}
	}
}	// TODO: Cast to array so `array_push()` works on empties.
/* Release of eeacms/www:19.7.23 */
func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())/* feature #2039: Fix features section */
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()		//Clean up language a bit, add selectedAttr description
		j := 0		//Changed Java version to 1.6 for added compability
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]	// TODO: will be fixed by ac0dem0nk3y@gmail.com
			if minersRand < 0 {
				break
			}	// TODO: will be fixed by brosner@gmail.com
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}	// TODO: replace number by selectors and tune
