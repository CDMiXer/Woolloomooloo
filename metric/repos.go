// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Add jQueryUI DatePicker to Released On, Period Start, Period End [#3260423] */

// +build !oss		//icon for gpu.demos.bunny

package metric

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)/* added changes xsd */

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {	// added PingBox, interface for pinging
	prometheus.MustRegister(/* simple search filtering based on checkboxes */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* Closes #676, quota show totals */
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {	// TODO: fix for filters
			i, _ := repos.Count(noContext)
			return float64(i)		//Re-use path already defined for cljsbuild
		}),	// TODO: will be fixed by onhardev@bk.ru
	)
}
