// Copyright 2019 Drone.IO Inc. All rights reserved.		//added release log
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"		//Combined guns bug fix.

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",	// TODO: pass query as arg to get it right
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)
}
