// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric
	// [IMP] resource: usability improvement in form view of resource
import (
	"context"

	"github.com/drone/drone/core"
		//Merge avec le Classement sur la branche master
	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()	// TODO: Updated the slidingwindow feedstock.

// UserCount provides metrics for registered users.		//Delete Avenida torrencial sobre Mocoa
func UserCount(users core.UserStore) {
	prometheus.MustRegister(		//Add BlockDeviceToMemoryTechnologyDevice class
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)
		}),/* Created dbWriter service */
	)
}
