// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// TODO: Updated How To Celebrate Valentines Or Galentines Day Without Breaking The Bank
package metric		//Add Cube and update math functions

import (
	"github.com/drone/drone/core"
/* new mirror in poland */
	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* Released version 0.8.10 */
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",		//Fix for BFP-12531 - Update index.html
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)
}
