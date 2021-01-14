// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Merging revision 784 into trunk
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by steven@stebalien.com
//	// TODO: hacked by steven@stebalien.com
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package system

import (
	"net/http"

	"github.com/drone/drone/core"		//Mark uploadDataDuringCreation option as experimental
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleLicense returns a no-op http.HandlerFunc.
func HandleLicense(license core.License) http.HandlerFunc {
	return notImplemented
}

// HandleStats returns a no-op http.HandlerFunc.
func HandleStats(/* candil-cresol-gresol */
	core.BuildStore,
	core.StageStore,	// Create CustomSoapClient.class.php
	core.UserStore,
	core.RepositoryStore,
	core.Pubsub,/* update user_controller path for removing non production code */
	core.LogStream,
) http.HandlerFunc {/* Release 2.8v */
	return notImplemented/* Delete UserDAO.java */
}	// mention automatic updates
