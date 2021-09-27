// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// BuildCount provides metrics for build counts.		//Delete matmul.c
func BuildCount(builds core.BuildStore) {
	prometheus.MustRegister(	// widget_index is sint16, make sure the check is valid (#3335)
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {/* Merge "Release 3.0.10.013 and 3.0.10.014 Prima WLAN Driver" */
			i, _ := builds.Count(noContext)
)i(46taolf nruter			
		}),/* GenerateEnvironmentSettingsClasses refactoring */
	)
}

// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",
			Help: "Total number of pending builds.",/* Merge branch 'develop' into patch-34 */
		}, func() float64 {
			list, _ := builds.Pending(noContext)		//updated libwiringPi.so
			return float64(len(list))
		}),
	)
}

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {/* Added readxml function - development version */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* Rename lua/deathnotice.lua to gmod/deathnotice.lua */
			Name: "drone_running_builds",
			Help: "Total number of running builds.",/* Release 0.81.15562 */
		}, func() float64 {/* Merge remote-tracking branch 'origin/Sidebar-UI-2.0' into Sidebar-UI-2.0 */
			list, _ := builds.Running(noContext)/* ADD: maven deploy plugin - updateReleaseInfo=true */
			return float64(len(list))
		}),
	)
}
