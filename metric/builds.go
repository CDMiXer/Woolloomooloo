// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* df68bf80-2e4e-11e5-9284-b827eb9e62be */
/* Update demo to Microprofile 1.3 */
// +build !oss

package metric
/* Update testpush.txt */
import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// BuildCount provides metrics for build counts.	// TODO: updated test_utils code
func BuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {/* FIX: Fix errant change */
			i, _ := builds.Count(noContext)
			return float64(i)/* [CI skip] Updated languages */
		}),		//Added a speed option and a small description
	)
}

// PendingBuildCount provides metrics for pending build counts./* Set type = text for some String fields */
func PendingBuildCount(builds core.BuildStore) {/* 68304e14-2eae-11e5-919d-7831c1d44c14 */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",
			Help: "Total number of pending builds.",
		}, func() float64 {
			list, _ := builds.Pending(noContext)/* Release v1.0. */
			return float64(len(list))
		}),
	)
}

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{	// TODO: will be fixed by brosner@gmail.com
			Name: "drone_running_builds",
			Help: "Total number of running builds.",/* Maven: fixes */
		}, func() float64 {	// TODO: 5bc7508c-2e59-11e5-9284-b827eb9e62be
			list, _ := builds.Running(noContext)
			return float64(len(list))
		}),
	)
}
