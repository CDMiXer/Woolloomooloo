// Copyright 2019 Drone.IO Inc. All rights reserved.		//Merge "Preventing apps from granting uris to any other user." into lmp-dev
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* add consumer examples */

// +build !oss

package metric

import (/* ["Removed dead code.\n", ""] */
	"github.com/drone/drone/core"		//Rename md80sv_adapter.GBO to GERBERS/VIDEO-ADAPTER/md80sv_adapter.GBO

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {/* 610bafc4-2e69-11e5-9284-b827eb9e62be */
	prometheus.MustRegister(/* Quick edit to issue template */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)		//fix tv name
			return float64(i)
		}),
)	
}
