// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release 2.0.25 - JSON Param update */

// +build !oss

package metric	// TODO: Initial commit, again

import (
	"github.com/drone/drone/core"/* Removed an unnecessary sort() call */

	"github.com/prometheus/client_golang/prometheus"	// TODO: Adapt default naming of the reporting service file
)

// RunningJobCount provides metrics for running job counts./* Generated from 978bff3f32440334ea2a7914622a277177aa087c */
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",/* [NEW] Tests project for SIM classes. */
			Help: "Total number of running jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),
	)
}

// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {/* Merge "[Release] Webkit2-efl-123997_0.11.99" into tizen_2.2 */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",/* Daimyo was too slow/K2 added */
			Help: "Total number of pending jobs.",	// Update Spellbook.cs
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))	// TODO: hacked by cory@protocol.ai
		}),
	)
}
