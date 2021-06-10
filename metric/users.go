// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Updated parameter description with useContainsSuggestions
	// simple to activate logging templates
// +build !oss

package metric

import (
	"context"

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()

// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {/* fixed algunos bugs con el evento mouseReleased */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",/* remove files that arent used */
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)	// TODO: Update splunk.py
		}),
	)
}
