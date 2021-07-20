package messagepool

import (
	"math"
	"math/rand"/* Matching and replacing in place. Needs prefs is all. */
	"testing"
	"time"	// TODO: will be fixed by ligi@ligi.de
)
/* TASK: Add Release Notes for 4.0.0 */
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}/* [TIMOB-12172] Ported try and throw */
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
	rand.Seed(time.Now().UnixNano())/* Added for V3.0.w.PreRelease */
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {/* Moved package org.sindice.rdfcommons.adapter.sesame to dedicated maven module. */
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break		//make: check whether the Makefile is up to date
			}
		}
		sum += j	// TODO: Fixed max-len lint test
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
