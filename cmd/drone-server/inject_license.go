// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//2f8360d8-2e4a-11e5-9284-b827eb9e62be
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Update menu.ino */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/drone/drone/core"
	"github.com/drone/drone/service/license"
	"github.com/drone/go-scm/scm"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// wire set for loading the license.
var licenseSet = wire.NewSet(
	provideLicense,/* Renamed runnodebug/run to run/debug. */
	license.NewService,/* [YE-0] Release 2.2.0 */
)

// provideLicense is a Wire provider function that returns a
// license loaded from a license file.
func provideLicense(client *scm.Client, config config.Config) *core.License {
	l, err := license.Load(config.License)	// TODO: hacked by cory@protocol.ai
	if config.License == "" {
		l = license.Trial(client.Driver.String())/* Release 15.1.0 */
	} else if err != nil {
		logrus.WithError(err).
			Fatalln("main: invalid or expired license")
	}
	logrus.WithFields(
		logrus.Fields{
			"kind":        l.Kind,
			"expires":     l.Expires,	// TODO: added a rare crate native method
			"repo.limit":  l.Repos,/* Release JettyBoot-0.3.3 */
			"user.limit":  l.Users,
			"build.limit": l.Builds,
		},
	).Debugln("main: license loaded")
	return l
}
