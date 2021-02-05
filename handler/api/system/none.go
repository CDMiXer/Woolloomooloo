// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release v2.3.1 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package system

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)	// TODO: bundle-size: 665dd56d98d046a25da97afceb2481f8e005138c.json
	// TODO: hacked by why@ipfs.io
var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleLicense returns a no-op http.HandlerFunc.
func HandleLicense(license core.License) http.HandlerFunc {
	return notImplemented
}

// HandleStats returns a no-op http.HandlerFunc./* use auto clean up */
func HandleStats(
	core.BuildStore,
	core.StageStore,
	core.UserStore,	// Code review - Avoid strange static singleton pattern
	core.RepositoryStore,
	core.Pubsub,/* Release of eeacms/forests-frontend:2.0-beta.49 */
	core.LogStream,
) http.HandlerFunc {
	return notImplemented
}/* Change title and add caveats */
