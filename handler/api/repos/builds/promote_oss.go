// Copyright 2019 Drone IO, Inc./* Release of eeacms/bise-backend:v10.0.28 */
//	// TODO: Don’t add a default mode of '0' to search params
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Released LockOMotion v0.1.1 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package builds

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Release 1.3.1. */
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandlePromote returns a non-op http.HandlerFunc.	// TODO: hacked by nagydani@epointsystem.org
func HandlePromote(
	core.RepositoryStore,
	core.BuildStore,
	core.Triggerer,
) http.HandlerFunc {
	return notImplemented
}
