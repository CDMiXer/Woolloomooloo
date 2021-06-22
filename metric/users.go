// Copyright 2019 Drone.IO Inc. All rights reserved./* Rename Networking/GetHost/gethost.py to Network/GetHost/gethost.py */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* #30 - Release version 1.3.0.RC1. */
// +build !oss

package metric

import (
	"context"/* Update README with intentions. */

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()
/* Fix create download page. Release 0.4.1. */
// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",	// TODO: hacked by igor@soramitsu.co.jp
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)
		}),
	)
}
