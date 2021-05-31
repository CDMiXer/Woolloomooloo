package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {	// Adding MySQL Driver
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)		//Add a reference to the multipart file uploader from commons-fileupload.
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out/* Release version 0.3.7 */
	})
	return noWinnersProbCache
}	// TODO: hacked by jon@atack.com

var noWinnersProbAssumingCache []float64/* include sms shortcodes on wall */
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)		//grub-rescue-pc.postinst: Build USB rescue image.
			result := math.Exp((math.Log(Mu) * x) - lg - cond)		//Merge "ARM: dts: msm: Add support for voice svc driver"
			return result
		}	// TODO: Merge branch 'master' into quick-styles

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()
	}	// TODO: will be fixed by julia@jvns.ca
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--
	}	// TODO: hacked by peterke@gmail.com
	return r
}
	// Daily work, making it useful for the toyDB. First commit use_minimal.py
func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq/* Add logger to media-alchemyst */
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}
		if p == 0 {	// TODO: hacked by souzau@yandex.com
			if x == 0 {
				return 1.0
			}
			return 0.0/* Release 1.2.0, closes #40 */
		}
		if p == 1 {		//add webrat style matchers
			if x == trials {
				return 1.0
			}/* f705f33e-2e5c-11e5-9284-b827eb9e62be */
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
