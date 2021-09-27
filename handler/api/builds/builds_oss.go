// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Fixed links and edited ocntent
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* F1DQae0oKKvcHfIGpzs54W7iEaFhrRcN */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release notes for 0.9.17 (and 0.9.16). */

// +build oss	// f9b69ea0-2e4b-11e5-9284-b827eb9e62be

package builds

import (
	"net/http"
/* Release native object for credentials */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleIncomplete returns a no-op http.HandlerFunc.
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {
	return notImplemented
}
