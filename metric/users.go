// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (/* Merge "Add Octavia charm" */
	"context"

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"	// TODO: hacked by mail@overlisted.net
)

var noContext = context.Background()

// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{		//Students divided by active and inactive (course)
			Name: "drone_user_count",	// TODO: updated to read in lower triangle matices
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)	// TODO: hacked by hugomrdias@gmail.com
			return float64(i)
		}),
	)		//add description for actions
}
