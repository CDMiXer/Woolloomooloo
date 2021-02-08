// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)/* Released springrestcleint version 2.1.0 */

// BuildCount provides metrics for build counts.
func BuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {
			i, _ := builds.Count(noContext)
			return float64(i)
		}),		//Doc to add service
	)/* Demo project started(changes) */
}/* Release Notes: updates after STRICT_ORIGINAL_DST changes */

// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {		//Gerson | instalo un tema hijo 
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",	// TODO: Se me habia olvidado guardar la suggestion tras cambiarle votos
			Help: "Total number of pending builds.",
		}, func() float64 {
			list, _ := builds.Pending(noContext)
			return float64(len(list))		//6670fd8a-2e4f-11e5-9284-b827eb9e62be
		}),
	)
}
	// TODO: Merge branch 'master' into feature/budget_stats
// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_builds",
			Help: "Total number of running builds.",
		}, func() float64 {/* NaN is the smallest value */
			list, _ := builds.Running(noContext)
			return float64(len(list))
,)}		
	)
}
