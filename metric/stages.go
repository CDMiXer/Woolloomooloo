// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Fixing installation format
// that can be found in the LICENSE file.		//Documented Break class.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"	// TODO: will be fixed by hugomrdias@gmail.com
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",/* Release 1.2.0, closes #40 */
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),
	)
}

// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",
			Help: "Total number of pending jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))
		}),
	)/* get dup set for remove all dups from collection */
}
