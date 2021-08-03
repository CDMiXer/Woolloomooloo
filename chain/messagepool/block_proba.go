package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64		//db/upnp/Discovery: use monotonic clock instead of time()
var noWinnersProbOnce sync.Once
		//spin position
func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})/* Samples: DynTex - can be handled by RTSS, no need for custom shaders */
	return noWinnersProbCache/* 0.17.2: Maintenance Release (close #30) */
}

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {		//Fix: Missing jquery.mCustomScrollbar.min.css
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5/* Eggdrop v1.8.2 Release Candidate 2 */
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)	// TODO: Sửa lỗi cảnh báo	
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

func binomialCoefficient(n, k float64) float64 {		//Better detection of bvh cache file permission issue
	if k > n {
		return math.NaN()/* update error slack<>mattermost */
	}
	r := 1.0		//First commit, just testing features and building base libraries
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
		if x > trials {/* Updated Readme for 4.0 Release Candidate 1 */
			return 0/* Release v0.21.0-M6 */
		}
		if p == 0 {/* Create play.md */
			if x == 0 {
				return 1.0	// Add ability to delete individual resource.
			}		//Weng mit Stanford geflirtet
			return 0.0
		}		//try installing llvm-5.0-dev
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
