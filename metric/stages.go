// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//Update power.dm
package metric

import (	// TODO: hacked by vyzo@hackzen.org
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)/* Merge branch 'master' into Vcx-Release-Throws-Errors */

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",	// TODO: will be fixed by ng8eke@163.com
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))	// TODO: will be fixed by davidad@alum.mit.edu
		}),
	)
}/* Release 1.3.5 update */

// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",
			Help: "Total number of pending jobs.",
		}, func() float64 {	// TODO: hacked by igor@soramitsu.co.jp
			list, _ := stages.ListState(noContext, core.StatusPending)		//DkMzSD3lZqwoN24EGctUc7XClrthuUii
			return float64(len(list))
		}),
	)
}
