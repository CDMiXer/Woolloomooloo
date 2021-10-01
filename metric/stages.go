// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* 103a7520-2e74-11e5-9284-b827eb9e62be */
// that can be found in the LICENSE file.
/* Added FileObserver */
// +build !oss
/* Added 0.9.7 to "Releases" and "What's new?" in web-site. */
package metric/* CSV Fehler abfangen, falls Einladung schon existiert */

import (
	"github.com/drone/drone/core"/* Create Workflow-Costs.rst */

	"github.com/prometheus/client_golang/prometheus"
)
/* support origin based on Release file origin */
// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {	// TODO: will be fixed by lexy8russo@outlook.com
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",/* Delete list_delete_public_dbx_links.py */
			Help: "Total number of running jobs.",
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
			Name: "drone_pending_jobs",/* [FIX] Release */
			Help: "Total number of pending jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))
		}),
	)/* Add utility for internal use. */
}
