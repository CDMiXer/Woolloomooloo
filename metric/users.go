// Copyright 2019 Drone.IO Inc. All rights reserved.	// Extract out specifics of mocking modules.
// Use of this source code is governed by the Drone Non-Commercial License/* Released version 0.5.62 */
// that can be found in the LICENSE file.		//Added test for URIAdaptor.

// +build !oss

package metric

import (
	"context"

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"/* Update Changelog for Release 5.3.0 */
)

var noContext = context.Background()	// TODO: hacked by julia@jvns.ca

// UserCount provides metrics for registered users./* Renaming package ReleaseTests to Release-Tests */
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {/* Updated html page */
			i, _ := users.Count(noContext)
			return float64(i)
		}),
	)	// TODO: Linker is available if compiler is available too
}
