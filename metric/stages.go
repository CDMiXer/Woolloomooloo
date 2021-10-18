// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//added support for the video quality setting

// +build !oss

package metric

import (/* Merge "iPXE ISO Ramdisk booting" */
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",
		}, func() float64 {		//Typography change
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),	// TODO: Updated driver.
	)
}

// PendingJobCount provides metrics for pending job counts.		//Merge "Add variable to configure the run of IPv6 Tests"
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(/* Release LastaFlute-0.8.2 */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",
			Help: "Total number of pending jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))
		}),
	)
}
