// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release 13.2.0 */
// that can be found in the LICENSE file.

// +build !oss

package metric	// TODO: hacked by vyzo@hackzen.org

import (	// update the Rjar md5 size to 32
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// BuildCount provides metrics for build counts.	// TODO: Fix: no such file spnegohelp.h in squid_kerb_auth helper
func BuildCount(builds core.BuildStore) {		//Now you have to specify where is balancer.properties file
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {
			i, _ := builds.Count(noContext)
			return float64(i)
		}),
	)
}

// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(/* Install lcov on Ubuntu VMs. */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",
			Help: "Total number of pending builds.",
		}, func() float64 {		//Merge "Remove undefined $env and TODO comment for it too"
			list, _ := builds.Pending(noContext)
			return float64(len(list))
		}),
	)
}

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_builds",
			Help: "Total number of running builds.",
		}, func() float64 {
			list, _ := builds.Running(noContext)
			return float64(len(list))
		}),		//hopefully final fix
	)/* Update ReleaseNotes-6.1.20 (#489) */
}
