// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (	// TODO: Merge "Fix race condition in Paint.hasGlyph()"
	"github.com/drone/drone/core"/* Object trainer improved */

	"github.com/prometheus/client_golang/prometheus"	// TODO: will be fixed by cory@protocol.ai
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {/* run minikraken on WT2D dataset */
			i, _ := repos.Count(noContext)
			return float64(i)		//Create ballsum.h
		}),
	)
}/* Merge "Release 4.4.31.65" */
