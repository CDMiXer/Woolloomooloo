// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Merge "Add simple script to help testing image membership" */
package metric

import (/* Merge branch 'master' into create-access-key-fix */
	"github.com/drone/drone/core"
	// TODO: Merge "Move resource doc generation to doc/source/ext"
	"github.com/prometheus/client_golang/prometheus"
)
/* Deleted CtrlApp_2.0.5/Release/TestClient.obj */
// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {/* Release: Making ready for next release cycle 4.5.3 */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",	// TODO: content tweaks.
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),/* Release Checklist > Bugs List  */
	)
}
