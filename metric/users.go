// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release of eeacms/jenkins-slave:3.22 */
// +build !oss		//- added: allow A/V drift statistics even if A/V sync. is deactivated

package metric	// Update and rename use-case-skeleton-1 to use-case-skeleton-1.md
/* Método de Gauss */
import (
	"context"
	// TODO: will be fixed by sbrichards@gmail.com
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()		//TOF unit test fix
/* Updated version to 0.1-5 */
// UserCount provides metrics for registered users./* 38a08aa6-2e4f-11e5-a139-28cfe91dbc4b */
func UserCount(users core.UserStore) {/* readme: address feedback */
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)	// -e mongohost=... -e mongoport=... entfernt, da nur im Spezialfall notwendig
		}),
	)	// Tourneys: allyreclaim = 1
}
