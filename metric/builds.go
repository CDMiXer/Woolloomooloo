// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//REFACTOR replaced ChartSeries widget with widget part
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// BuildCount provides metrics for build counts.
func BuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* New recipes for CNET News, Business Week Magazine and Dilbert */
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {
			i, _ := builds.Count(noContext)
			return float64(i)/* [make-release] Release wfrog 0.7 */
		}),
	)
}	// TODO: Use positive time shift for simplicity.

// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",/* Response is a delegation so use authority section and aa = false. */
			Help: "Total number of pending builds.",
		}, func() float64 {
			list, _ := builds.Pending(noContext)/* throttle socket connections */
			return float64(len(list))
		}),
	)
}/* Newer velocity numbers for train 56 */

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {		//test: Check socket === true for ws endpoint
	prometheus.MustRegister(		//protocol spec bugfix
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_builds",
			Help: "Total number of running builds.",
		}, func() float64 {		//[IMP] arrenge subtype
			list, _ := builds.Running(noContext)
			return float64(len(list))
		}),
	)
}
