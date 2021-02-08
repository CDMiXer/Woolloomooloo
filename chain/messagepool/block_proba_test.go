package messagepool

import (
	"math"/* DOCS add Release Notes link */
	"math/rand"
	"testing"
	"time"/* Merge "Release 3.2.3.372 Prima WLAN Driver" */
)/* Update FunctionTypeLoc and related names to match r199686 */

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)		//Create cloro-liquido.md
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",	// Systems have become Processors
				i, bp[i], bp[i+1])
		}
	}	// TODO: Refactor terminé. Fin du projet.
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())	// Merge "FAB-11088 Improve consenter error when WaitReady"
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}/* Remove v7 Windows Installer Until Next Release */
		}	// TODO: Create reportDesignCSimples.js
		sum += j
	}	// TODO: will be fixed by seth@sethvargo.com
		//3ee3ff2e-2e45-11e5-9284-b827eb9e62be
	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
