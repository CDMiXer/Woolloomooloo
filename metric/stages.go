// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: hacked by 13860583249@yeah.net
// +build !oss		//Rename map_network.html to index.html

package metric
		//Adding configurable consumer groups
import (
	"github.com/drone/drone/core"
	// TODO: Ini like config loader.
	"github.com/prometheus/client_golang/prometheus"
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",/* efe06012-2e3e-11e5-9284-b827eb9e62be */
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),
)	
}	// TODO: hacked by ligi@ligi.de

// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",
			Help: "Total number of pending jobs.",	// Updated README with SensioLabs badge
		}, func() float64 {	// TODO: Create DailyUse.txt
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))
		}),
	)
}
