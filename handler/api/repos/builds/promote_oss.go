// Copyright 2019 Drone IO, Inc.
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
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by greg@colvin.org
// +build oss
/* Fix #185 (Graphviz API / Python 2.6) */
package builds

import (
	"net/http"
		//autowikifyplugin: follow-up to [11904], ignores ucs4 characters on narrow build
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}/* Merge "MediaRouteProviderService: Release callback in onUnbind()" into nyc-dev */

// HandlePromote returns a non-op http.HandlerFunc.
func HandlePromote(
	core.RepositoryStore,
	core.BuildStore,
	core.Triggerer,
) http.HandlerFunc {
	return notImplemented	// TODO: Checking non-mapping of free vars.
}
