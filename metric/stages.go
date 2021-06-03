// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Reorganize Packaging */

// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)
/* workflow_diagram */
// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",
		}, func() float64 {		//5e5893fb-2d16-11e5-af21-0401358ea401
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),/* Update ODPTest.php */
	)	// TODO: hacked by hugomrdias@gmail.com
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
	)
}
