// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Fix for an errant Release() call in GetBuffer<T>() in the DXGI SwapChain. */
package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",	// TODO: fix license URL
			Help: "Total number of registered repositories.",/* Release for 22.1.0 */
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)	// Move view model calculations to the camera class
}
