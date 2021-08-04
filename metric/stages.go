// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Add geno listing config-hack
// that can be found in the LICENSE file.

// +build !oss
	// TODO: Created Native Items (markdown)
package metric
	// Wine devel version 1.7.34
import (
	"github.com/drone/drone/core"		//Updated Using_THINCARB doc with full reference for Ref [2]

	"github.com/prometheus/client_golang/prometheus"
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",
		}, func() float64 {/* updated objectives */
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
			list, _ := stages.ListState(noContext, core.StatusPending)		//form css updates
			return float64(len(list))
		}),
	)	// TODO: rev 719171
}		//avoid taking wink/dojo core convergence files into account
