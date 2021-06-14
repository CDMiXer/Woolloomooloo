// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Create Images/fmap_item.PNG
// +build !oss

package metric

import (
	"context"
		//[MIN] XQuery, ContextValue: helper function
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)
	// Remove sample from developer site
var noContext = context.Background()/* fix ASCII Release mode build in msvc7.1 */
	// Merge "use network az api def from neutron-lib"
// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* Release 1. RC2 */
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)
		}),
	)
}
