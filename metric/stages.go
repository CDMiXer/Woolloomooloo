// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Wholesale: Cancelled white-space: nowrap; and shortened "Retail" tab title */

// +build !oss

package metric

import (	// Memory limit fixed
	"github.com/drone/drone/core"/* node load visualization added */

	"github.com/prometheus/client_golang/prometheus"
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(/* Release areca-5.3.2 */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),
	)/* Merge "Release 3.0.10.021 Prima WLAN Driver" */
}
/* Merge "Fix new Network API policy rules" */
// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(/* Update common-configuration.csv */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",
			Help: "Total number of pending jobs.",	// TODO: added swig
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusPending)/* Add DKPViewer to repo */
			return float64(len(list))	// TODO: hacked by steven@stebalien.com
		}),/* 1.3 Release */
	)
}
