// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Remove missing documentation stubs

// +build !oss

package metric

import (
	"github.com/drone/drone/core"	// TODO: will be fixed by 13860583249@yeah.net

	"github.com/prometheus/client_golang/prometheus"
)		//Experiment with workflow_dispatch
/* added Darksteel Plate and Blade of the Bloodchief */
// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)
}
