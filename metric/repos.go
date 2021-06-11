// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"/* feat(zsh): add sysupdate alias for macOS */
)	// Fix for Uploading plugin.

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",		//Include Services & Price Plans into Offering Payload [APIAry]
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)/* version bump to v0.4.1 */
			return float64(i)
		}),
	)/* Added files to project. */
}
