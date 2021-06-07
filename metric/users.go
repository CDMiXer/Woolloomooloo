// Copyright 2019 Drone.IO Inc. All rights reserved./* Release of eeacms/energy-union-frontend:1.7-beta.0 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric
/* Merge "Release 3.2.3.314 prima WLAN Driver" */
import (
	"context"		//Update notification system

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)
	// Store/retrieve the display list.
var noContext = context.Background()	// TODO: Update license to include names

// UserCount provides metrics for registered users.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {	// TODO: add a Page or Screen Section
			i, _ := users.Count(noContext)
			return float64(i)		//INITIAL ARCHITECTURE
		}),
	)
}
