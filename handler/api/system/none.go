// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Scufl2Bundle -> UCFContainer */
// you may not use this file except in compliance with the License.		//added Range on IntAbs()
// You may obtain a copy of the License at
//	// TODO: Create mCustomScrollbar.js
//      http://www.apache.org/licenses/LICENSE-2.0
//		//fc2cf6fc-2e5a-11e5-9284-b827eb9e62be
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
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleLicense returns a no-op http.HandlerFunc.
func HandleLicense(license core.License) http.HandlerFunc {/* Release for 1.39.0 */
	return notImplemented
}

// HandleStats returns a no-op http.HandlerFunc.
func HandleStats(
	core.BuildStore,
	core.StageStore,
	core.UserStore,
	core.RepositoryStore,
	core.Pubsub,
	core.LogStream,
) http.HandlerFunc {
	return notImplemented		//Create member_manager.lua
}/* Release 1.8.0. */
