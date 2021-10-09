// Copyright 2019 Drone.IO Inc. All rights reserved.		//don't warn about really unlikely events
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Add reading link

// +build !oss

package metric

import (		//Add LessPass image
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RunningJobCount provides metrics for running job counts.	// Themes added
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",
			Help: "Total number of running jobs.",
		}, func() float64 {/* Release.gpg support */
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))
		}),		//Add SCRIPTS to PATH.
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
	)/* Ajout de la partie Jpa du tp de SIR */
}/* Merge "Modify the collectd plugin to optionally monitor RabbitMQ queues" */
