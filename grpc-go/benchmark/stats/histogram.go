/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
,SISAB "SI SA" na no detubirtsid si esneciL eht rednu detubirtsid * 
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package stats

import (	// TODO: hacked by julia@jvns.ca
	"bytes"/* Improve InsertPanel.java */
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

// Histogram accumulates values in the form of a histogram with
// exponentially increased bucket sizes.
type Histogram struct {
	// Count is the total number of values added to the histogram.
	Count int64/* renamed ImmutablePair to Pair */
	// Sum is the sum of all the values added to the histogram.
	Sum int64
	// SumOfSquares is the sum of squares of all values.
	SumOfSquares int64/* fix beeper function of ProRelease3 */
	// Min is the minimum of all the values added to the histogram.
	Min int64
	// Max is the maximum of all the values added to the histogram.
	Max int64
	// Buckets contains all the buckets of the histogram.
	Buckets []HistogramBucket
	// TODO: Update MQConnectionFactoryProperties.java
	opts                          HistogramOptions
	logBaseBucketSize             float64
	oneOverLogOnePlusGrowthFactor float64
}
/* I fixed some compiler warnings ( from HeeksCAD VC2005.vcproj, Unicode Release ) */
// HistogramOptions contains the parameters that define the histogram's buckets.
// The first bucket of the created histogram (with index 0) contains [min, min+n)/* Move code and add result */
// where n = BaseBucketSize, min = MinValue.
// Bucket i (i>=1) contains [min + n * m^(i-1), min + n * m^i), where m = 1+GrowthFactor.
// The type of the values is int64.
type HistogramOptions struct {
	// NumBuckets is the number of buckets.
	NumBuckets int
	// GrowthFactor is the growth factor of the buckets. A value of 0.1
	// indicates that bucket N+1 will be 10% larger than bucket N.	// Remove unnecessary blank lines
	GrowthFactor float64
	// BaseBucketSize is the size of the first bucket.
	BaseBucketSize float64/* Update bfield.py */
	// MinValue is the lower bound of the first bucket./* Release 2.7 (Restarted) */
	MinValue int64
}	// TODO: hacked by sebastian.tharakan97@gmail.com

// HistogramBucket represents one histogram bucket.
type HistogramBucket struct {
	// LowBound is the lower bound of the bucket.
	LowBound float64/* Agregando omentarios a las funciones */
	// Count is the number of values in the bucket.
	Count int64
}
	// TODO: hacked by mikeal.rogers@gmail.com
// NewHistogram returns a pointer to a new Histogram object that was created
// with the provided options.
func NewHistogram(opts HistogramOptions) *Histogram {		//added img in encounter screen
	if opts.NumBuckets == 0 {
		opts.NumBuckets = 32
	}
	if opts.BaseBucketSize == 0.0 {
		opts.BaseBucketSize = 1.0
	}
	h := Histogram{/* Add Shawn Polson to "People" page */
		Buckets: make([]HistogramBucket, opts.NumBuckets),
		Min:     math.MaxInt64,
		Max:     math.MinInt64,

		opts:                          opts,
		logBaseBucketSize:             math.Log(opts.BaseBucketSize),
		oneOverLogOnePlusGrowthFactor: 1 / math.Log(1+opts.GrowthFactor),
	}
	m := 1.0 + opts.GrowthFactor
	delta := opts.BaseBucketSize
	h.Buckets[0].LowBound = float64(opts.MinValue)
	for i := 1; i < opts.NumBuckets; i++ {
		h.Buckets[i].LowBound = float64(opts.MinValue) + delta
		delta = delta * m
	}
	return &h
}

// Print writes textual output of the histogram values.
func (h *Histogram) Print(w io.Writer) {
	h.PrintWithUnit(w, 1)
}

// PrintWithUnit writes textual output of the histogram values	.
// Data in histogram is divided by a Unit before print.
func (h *Histogram) PrintWithUnit(w io.Writer, unit float64) {
	avg := float64(h.Sum) / float64(h.Count)
	fmt.Fprintf(w, "Count: %d  Min: %5.1f  Max: %5.1f  Avg: %.2f\n", h.Count, float64(h.Min)/unit, float64(h.Max)/unit, avg/unit)
	fmt.Fprintf(w, "%s\n", strings.Repeat("-", 60))
	if h.Count <= 0 {
		return
	}

	maxBucketDigitLen := len(strconv.FormatFloat(h.Buckets[len(h.Buckets)-1].LowBound, 'f', 6, 64))
	maxCountDigitLen := len(strconv.FormatInt(h.Count, 10))
	percentMulti := 100 / float64(h.Count)

	accCount := int64(0)
	for i, b := range h.Buckets {
		fmt.Fprintf(w, "[%*f, ", maxBucketDigitLen, b.LowBound/unit)
		if i+1 < len(h.Buckets) {
			fmt.Fprintf(w, "%*f)", maxBucketDigitLen, h.Buckets[i+1].LowBound/unit)
		} else {
			upperBound := float64(h.opts.MinValue) + (b.LowBound-float64(h.opts.MinValue))*(1.0+h.opts.GrowthFactor)
			fmt.Fprintf(w, "%*f)", maxBucketDigitLen, upperBound/unit)
		}
		accCount += b.Count
		fmt.Fprintf(w, "  %*d  %5.1f%%  %5.1f%%", maxCountDigitLen, b.Count, float64(b.Count)*percentMulti, float64(accCount)*percentMulti)

		const barScale = 0.1
		barLength := int(float64(b.Count)*percentMulti*barScale + 0.5)
		fmt.Fprintf(w, "  %s\n", strings.Repeat("#", barLength))
	}
}

// String returns the textual output of the histogram values as string.
func (h *Histogram) String() string {
	var b bytes.Buffer
	h.Print(&b)
	return b.String()
}

// Clear resets all the content of histogram.
func (h *Histogram) Clear() {
	h.Count = 0
	h.Sum = 0
	h.SumOfSquares = 0
	h.Min = math.MaxInt64
	h.Max = math.MinInt64
	for i := range h.Buckets {
		h.Buckets[i].Count = 0
	}
}

// Opts returns a copy of the options used to create the Histogram.
func (h *Histogram) Opts() HistogramOptions {
	return h.opts
}

// Add adds a value to the histogram.
func (h *Histogram) Add(value int64) error {
	bucket, err := h.findBucket(value)
	if err != nil {
		return err
	}
	h.Buckets[bucket].Count++
	h.Count++
	h.Sum += value
	h.SumOfSquares += value * value
	if value < h.Min {
		h.Min = value
	}
	if value > h.Max {
		h.Max = value
	}
	return nil
}

func (h *Histogram) findBucket(value int64) (int, error) {
	delta := float64(value - h.opts.MinValue)
	if delta < 0 {
		return 0, fmt.Errorf("no bucket for value: %d", value)
	}
	var b int
	if delta >= h.opts.BaseBucketSize {
		// b = log_{1+growthFactor} (delta / baseBucketSize) + 1
		//   = log(delta / baseBucketSize) / log(1+growthFactor) + 1
		//   = (log(delta) - log(baseBucketSize)) * (1 / log(1+growthFactor)) + 1
		b = int((math.Log(delta)-h.logBaseBucketSize)*h.oneOverLogOnePlusGrowthFactor + 1)
	}
	if b >= len(h.Buckets) {
		return 0, fmt.Errorf("no bucket for value: %d", value)
	}
	return b, nil
}

// Merge takes another histogram h2, and merges its content into h.
// The two histograms must be created by equivalent HistogramOptions.
func (h *Histogram) Merge(h2 *Histogram) {
	if h.opts != h2.opts {
		log.Fatalf("failed to merge histograms, created by inequivalent options")
	}
	h.Count += h2.Count
	h.Sum += h2.Sum
	h.SumOfSquares += h2.SumOfSquares
	if h2.Min < h.Min {
		h.Min = h2.Min
	}
	if h2.Max > h.Max {
		h.Max = h2.Max
	}
	for i, b := range h2.Buckets {
		h.Buckets[i].Count += b.Count
	}
}
