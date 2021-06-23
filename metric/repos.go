// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Merge "Colorado Release note" */
// +build !oss/* Release v0.3 */

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)/* found an old screenshot when MozEmbed was used */

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{		//Merge branch 'dev_reddyice' into iseries
			Name: "drone_repo_count",/* Stats_template_added_to_ReleaseNotes_for_all_instances */
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)	// TODO: hacked by sjors@sprovoost.nl
			return float64(i)		//d161909c-2e3f-11e5-9284-b827eb9e62be
		}),
	)
}
