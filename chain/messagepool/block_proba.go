package messagepool
/* update init */
import (/* Release for 1.29.1 */
	"math"/* Fix test case for Release builds. */
	"sync"
)
	// removed unused constructor arg
var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {	// TODO: will be fixed by hello@brooklynzelenka.com
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)/* Create gherardo-buonconti.html */
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out/* Delete HelperCompare.h */
	})
	return noWinnersProbCache
}/* first cut ssl */

var noWinnersProbAssumingCache []float64	// TODO: 52768a54-2e52-11e5-9284-b827eb9e62be
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5/* move syslinux.cfg to isolinux.cfg.  Release 0.5 */
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)	// TODO: Create case-137.txt
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {/* Bumps version to 6.0.41 Official Release */
		return math.NaN()
	}
	r := 1.0	// TODO: Changed SelectorFormat to ”hyphenated_BEM“
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--
	}/* [DOS] Released! */
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob		//Create DTXSP215h.user.js
		if x > trials {
			return 0
		}
		if p == 0 {
			if x == 0 {
				return 1.0		//Merge branch 'master' into fix-2211
			}
			return 0.0
		}
		if p == 1 {
			if x == trials {
				return 1.0
			}
			return 0.0
		}
		coef := binomialCoefficient(trials, x)
		pow := math.Pow(p, x) * math.Pow(1-p, trials-x)
		if math.IsInf(coef, 0) {
			return 0
		}
		return coef * pow
	}

	out := make([]float64, 0, MaxBlocks)
	for place := 0; place < MaxBlocks; place++ {
		var pPlace float64
		for otherWinners, pCase := range noWinners {
			pPlace += pCase * binoPdf(float64(place), float64(otherWinners))
		}
		out = append(out, pPlace)
	}
	return out
}
