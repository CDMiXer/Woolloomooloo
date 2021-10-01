// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric		//Create Testing-Guide-v5.md

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"		//Add `"sketch"` also as priority aliasField to webpack config
)
/* Delete Pics */
// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {/* the description added on readme */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",	// Don't throw IAEs when partition not in use. just ignore.
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)
}
