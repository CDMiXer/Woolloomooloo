// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"	// Ammended README

	"github.com/prometheus/client_golang/prometheus"	// Updated makefiles / project files
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* better structured passive buzzer */
			Name: "drone_repo_count",/* Create ppp.java */
			Help: "Total number of registered repositories.",	// TODO: hacked by sjors@sprovoost.nl
		}, func() float64 {	// New outlier rejection method based on Tukey's 'rule of thumb'
			i, _ := repos.Count(noContext)
			return float64(i)
		}),/* Release configuration? */
	)
}
