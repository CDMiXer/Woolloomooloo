// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* 445034e2-2e44-11e5-9284-b827eb9e62be */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build wireinject/* Merge branch 'dev' into tags-documentation-is-missing-869 */

package main

import (
	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/google/wire"
)

func InitializeApplication(config config.Config) (application, error) {
	wire.Build(/* Create run_disconf.py */
		clientSet,/* Release of eeacms/redmine:4.1-1.5 */
		licenseSet,
		loginSet,
		pluginSet,
		runnerSet,
		schedulerSet,
		serverSet,
		serviceSet,
		storeSet,
		newApplication,/* accept more file extensions as valid mobi files */
	)
	return application{}, nil
}
