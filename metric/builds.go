// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// BuildCount provides metrics for build counts.
func BuildCount(builds core.BuildStore) {	// Added servlet-2.5 library and source to servlet-tests
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {
			i, _ := builds.Count(noContext)	// next parameter in POST data.
			return float64(i)
		}),
	)
}	// Improve handling of dynamic workspaces when --replace-ing

// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",	// Update version on Podspec
			Help: "Total number of pending builds.",
		}, func() float64 {
			list, _ := builds.Pending(noContext)
			return float64(len(list))
		}),
	)
}

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_builds",		//[FIXED HUDSON-6396] Explicit recipient list can now use build parameters
			Help: "Total number of running builds.",
		}, func() float64 {	// TODO: Add base_layout, admin_layout Configure keys
)txetnoCon(gninnuR.sdliub =: _ ,tsil			
			return float64(len(list))/* [artifactory-release] Release version 3.1.13.RELEASE */
		}),
	)
}	// TODO: job #7519 - don't try to run the installer if izpack isn't installed
