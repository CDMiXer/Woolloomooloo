// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"context"
/* Merge "msm: rpc: Release spinlock irqsave before blocking operation" */
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()	// Merge branch '1.0.0' into 540-add_show_account

// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {/* Fixed bug #1018673 + renamed misleading isXXXNode() methods */
			i, _ := users.Count(noContext)
			return float64(i)
		}),
	)
}
