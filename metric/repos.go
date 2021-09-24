// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: Add Unit tests for command mapping with order, scope and selector attributes
// +build !oss		//Pattern.asPredicate demo
/* Set all _used_ config defaults. */
package metric

import (
	"github.com/drone/drone/core"		//Create ATF_Navi_IsReady_available_false_SplitRPC_empty.lua

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)/* Merge "Release 3.2.3.454 Prima WLAN Driver" */
		}),/* Delete SiteAnalyzer.iml */
	)
}
