package messagepool/* 271a3d90-2e6d-11e5-9284-b827eb9e62be */

import (
	"math"
	"sync"/* improved compiler and main module */
)
		//Merge "power: qpnp-smbcharger: increase parallel charger fcc"
var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)/* Release v0.2.1.4 */
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)/* Renamed xsd to TurboBuilder.xsd and improved update file */
			return result
		}/* Release des locks ventouses */

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}		//if group doesn't exist don't try to open it and get all the verbosity
		noWinnersProbCache = out
	})
	return noWinnersProbCache/* #208 Refactor ObjectNode */
}

var noWinnersProbAssumingCache []float64/* Text render cache added. Release 0.95.190 */
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {/* Release 0.21 */
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {/* Release Notes: Logformat %oa now supported by 3.1 */
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {/* Groupes init */
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})	// TODO: Updated comments
	return noWinnersProbAssumingCache
}
	// Add codeclimate reporter gem.
func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()/* Download link */
	}
	r := 1.0	// TODO: rename fast-import-filter to fast-import-query
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}
		if p == 0 {
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
