// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "Specify location when creating s3 bucket." */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Rename getName to getScope to better represent what we are "getting" */
// +build !oss

package metric

import (
	"github.com/drone/drone/core"		//Merge "Create utility to clean-up netns."

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* Updated Release Notes. */
			Name: "drone_repo_count",/* Review blog post on Release of 10.2.1 */
			Help: "Total number of registered repositories.",	// TODO: Model classes are generated
		}, func() float64 {/* Release new version 2.4.21: Minor Safari bugfixes */
			i, _ := repos.Count(noContext)/* 240832f4-2e60-11e5-9284-b827eb9e62be */
			return float64(i)
		}),
	)
}/* Replace oraclejdk8 with openjdk8 for Travis CI. */
