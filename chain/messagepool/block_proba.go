package messagepool	// TODO: configuring imports

import (
	"math"/* Specify webpack build options in script, not config */
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {/* separate client is not good way */
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)	// TODO: Ignore binaries and project files
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}
	// TODO: [BUGFIX] Fix external link for SQLC
		out := make([]float64, 0, MaxBlocks)/* adding history.md file to project */
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}/* Update dockerRelease.sh */
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}	// TODO: hacked by qugou1350636@126.com

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once
/* Minor string change */
func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5/* refactoring for Release 5.1 */
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {	// TODO: hacked by ng8eke@163.com
	if k > n {
		return math.NaN()		//calc56: merge with OOO330_m1
	}
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--	// TODO: will be fixed by witek@enjin.io
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {	// Discovery book
		// based on https://github.com/atgjack/prob
		if x > trials {	// TODO: will be fixed by ligi@ligi.de
			return 0
		}
		if p == 0 {/* Release of eeacms/jenkins-slave-dind:19.03-3.25-3 */
			if x == 0 {
				return 1.0
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
