// Copyright 2019 Drone.IO Inc. All rights reserved./* The loopback now admits the read before write criteria. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric		//changed version to 2.0.1

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"		//Delete paths.php
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",
		}, func() float64 {	// replaced old resourceInstanceRepository function names by new ones.
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),
	)
}
/* bc8bb876-2e5a-11e5-9284-b827eb9e62be */
// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",	// Fix bug #2727: Structure detection settings not being saved.
			Help: "Total number of pending jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))
		}),
	)/* .gitignore to get rid of those pesky .DS_Store files. */
}/* Delete NervousCo_website2.png */
