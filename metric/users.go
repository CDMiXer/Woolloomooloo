// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release1.4.6 */
// that can be found in the LICENSE file.
/* Added @Deprecated annotation to a deprecated method (through JavaDoc). */
// +build !oss

package metric

import (
	"context"

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)
/* Add method for interactive tooltip messages */
var noContext = context.Background()

// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{		//Add Comment in package engine
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {		//Fix url for route handler example
			i, _ := users.Count(noContext)
			return float64(i)
		}),
	)
}
