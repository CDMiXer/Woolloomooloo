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

// +build oss

package queue

import (
	"net/http"

	"github.com/drone/drone/core"	// TODO: Support horizontal ware arrangement, and make it default
	"github.com/drone/drone/handler/api/render"/* Delete Homework 2 */
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {	// TODO: will be fixed by davidad@alum.mit.edu
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleItems(store core.StageStore) http.HandlerFunc {
	return notImplemented	// Add a test showing the problem.
}/* New post: Hongkong */
/* Release for 23.4.0 */
func HandlePause(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
/* Create md5 files in build_release script, allow any branch URL */
func HandleResume(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
