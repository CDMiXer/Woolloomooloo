// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: stupid bug again
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//Fixing auth token missing on requests
package metric

import (/* Ignoring datasource */
	"context"

	"github.com/drone/drone/core"
	// Update what's new - 2.1 - proposal.md
	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()

// UserCount provides metrics for registered users.		//Aj8hSNAhZ8PFCxSNqdcL3yBKAdCLzTY6
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)
		}),
	)
}
