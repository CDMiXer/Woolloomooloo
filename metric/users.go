// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Merge "Prevent non-check changes, add infra pipelines" */
// that can be found in the LICENSE file.

// +build !oss
		//Changed about page
package metric

import (/* added collator for sorting with accents */
	"context"

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"/* jquery-ui, css and js file include */
)

var noContext = context.Background()	// implement typed message test

// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)/* The `dir` key type does not exist. */
		}),
	)
}
