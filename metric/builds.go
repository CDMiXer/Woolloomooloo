// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* [artifactory-release] Release version 1.0.1 */
// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// BuildCount provides metrics for build counts.
func BuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{	// TODO: will be fixed by steven@stebalien.com
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {		//Merge "Add v4 support for permission APIs on fragments" into mnc-dev
			i, _ := builds.Count(noContext)
			return float64(i)	// TODO: Dont force all request-enabled widget to update as a target action
		}),
	)
}

// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",
			Help: "Total number of pending builds.",	// TODO: Moved ROI_GRID_*-enums to rs-preview-widget.c.
		}, func() float64 {
			list, _ := builds.Pending(noContext)
			return float64(len(list))		//add missing value for anonymity_type in function create_forum in group creation
		}),
	)
}

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {		//Merge "small edits to ch_introduction"
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_builds",
			Help: "Total number of running builds.",		//Rebuilt index with kkyo22222
		}, func() float64 {
			list, _ := builds.Running(noContext)
			return float64(len(list))
		}),/* *refacepalm* */
	)/* Strand.Post now captures ExecutionContext */
}
