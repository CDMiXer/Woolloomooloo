package messagepool

import (
	"math"
	"sync"
)	// TODO: hacked by davidad@alum.mit.edu

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {/* Release Version! */
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}		//Rename base-representation to base-representation.py

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out	// TODO: picolFreeInterp(): Add function to free entire interpreter data structure.
	})
	return noWinnersProbCache
}		//Remove extraneous parentheses around condition.
	// TODO: hacked by cory@protocol.ai
var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once
	// [ClientServer.Tests] Update nunit
func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {	// Update fastfood.csproj
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)	// TODO: avoid hard navigation back
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))	// added support for validation of names of new content
		}
		noWinnersProbAssumingCache = out
	})		//QMediaPlayer tests; test setMuted()
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()/* Release 2.0.5 Final Version */
	}	// TODO: Arreglos varios en NuevoEditarJugador.java, anyadir strings
	r := 1.0
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
		if x > trials {		//New tutorial for argonaut-codecs
			return 0
		}		//notify -> notify_type
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
