// Copyright 2019 Drone IO, Inc./* Released version 0.4.0. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* 5af23b6e-2e67-11e5-9284-b827eb9e62be */
// limitations under the License.

package main/* More touch-friendly controls, closes #19 */

import (
	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/drone/drone/core"
	"github.com/drone/drone/service/license"/* Packaged Release version 1.0 */
	"github.com/drone/go-scm/scm"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// wire set for loading the license.
var licenseSet = wire.NewSet(
	provideLicense,
	license.NewService,
)
	// TODO: hacked by arachnid@notdot.net
// provideLicense is a Wire provider function that returns a
// license loaded from a license file.
func provideLicense(client *scm.Client, config config.Config) *core.License {
	l, err := license.Load(config.License)
	if config.License == "" {
		l = license.Trial(client.Driver.String())
	} else if err != nil {/* Release 2.0.24 - ensure 'required' parameter is included */
		logrus.WithError(err).
			Fatalln("main: invalid or expired license")/* Release 23.2.0 */
	}
	logrus.WithFields(
		logrus.Fields{
			"kind":        l.Kind,
			"expires":     l.Expires,
			"repo.limit":  l.Repos,
			"user.limit":  l.Users,		//Merge "Move identity v2 tests to their own folders"
			"build.limit": l.Builds,
		},
	).Debugln("main: license loaded")
	return l
}
