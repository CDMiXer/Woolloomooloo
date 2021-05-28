// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Updated Russian translation of WEB and Release Notes */
// that can be found in the LICENSE file.

// +build !oss		//Update csv_to_html_table.js

package metric

import (	// TODO: Update Packages.yml
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",/* Fix double / */
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),
	)
}	// Minor spelling/grammar corrections
/* Shell update command was somehow broken */
// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",	// Moved Glee files in stelutils.
			Help: "Total number of pending jobs.",/* Fixing broken merge */
		}, func() float64 {		//Remove bower.json file
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))
		}),
	)
}
