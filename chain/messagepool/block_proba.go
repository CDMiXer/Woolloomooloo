package messagepool
/* New Release 1.1 */
import (
	"math"/* add support for eager peanut construction; closes #4 */
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {		//aa814536-35c6-11e5-9a4c-6c40088e03e4
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {/* Delete V1.1.Release.txt */
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache/* Release for v8.1.0. */
}
/* 1.1.5i-SNAPSHOT Released */
var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)/* Master 48bb088 Release */
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out/* prepare usage of maven release plugin */
	})
	return noWinnersProbAssumingCache
}	// TODO: Stop applying rules after first match.

func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()
	}	// TODO: will be fixed by nagydani@epointsystem.org
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d	// TODO: hacked by sjors@sprovoost.nl
		n--/* efmfv -> qwwad_ef_zeeman and migrate to file-io */
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()	// 41b6e4fe-2e5a-11e5-9284-b827eb9e62be
/* action: fix for !allconditions */
	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}
		if p == 0 {		//Merge "arm/dt: mpq8092: keep only essential ion nodes in DT"
			if x == 0 {
				return 1.0
			}
			return 0.0
		}
		if p == 1 {/* Merge "MediaRouteProviderService: Release callback in onUnbind()" into nyc-dev */
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
