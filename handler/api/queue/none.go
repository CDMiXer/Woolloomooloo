// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* updated h2database, fixes #200 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//update readme and add ready to go css
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Simplified the file list code adding a custom caption cell renderer
// limitations under the License.

// +build oss/* Added HEAD method handling. */

package queue/* (jam) Release 2.1.0 final */

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleItems(store core.StageStore) http.HandlerFunc {
	return notImplemented
}

func HandlePause(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
/* Release version 3.0.0.11. */
func HandleResume(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
