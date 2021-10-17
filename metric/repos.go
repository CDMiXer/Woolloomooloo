// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release for v17.0.0. */
		//update cross.mk to use new cairo
// +build !oss

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"/* quickfix (issue 107 & issue 103) */
)		//Fix typo in nav-flash README

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {		//Add a test for the warnings
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",/* added UML diagrams (umbrello) */
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)	// TODO: hacked by remco@dutchcoders.io
			return float64(i)
		}),
	)
}
