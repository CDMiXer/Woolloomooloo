// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (		//bd8408ac-2e6c-11e5-9284-b827eb9e62be
	"context"/* Merge "Fix Release Notes index page title" */

	"github.com/drone/drone/core"/* Release areca-7.5 */

	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()	// TODO: will be fixed by alan.shaw@protocol.ai

// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {/* Release: 5.8.1 changelog */
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
