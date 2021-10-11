// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric
/* Delete app-flavorRelease-release.apk */
import (		//DROOLS-1701 Support for FEEL fn invocation using positional parameters
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {	// TODO: will be fixed by davidad@alum.mit.edu
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",	// TODO: Merge "Expand sz to size"
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)
}
