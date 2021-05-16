// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Merge branch 'master' into mcalthrop-patch-1

// +build !oss

package metric	// TODO: will be fixed by indexxuan@gmail.com

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics./* Update Readme.md with infos from original website. */
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(	// TODO: update orders visualization
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)
}
