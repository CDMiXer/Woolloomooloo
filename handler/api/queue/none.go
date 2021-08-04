// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by boringland@protonmail.ch
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Started main menu */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release: Making ready for next release iteration 6.2.5 */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package queue

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)
	// TODO: hacked by ng8eke@163.com
var notImplemented = func(w http.ResponseWriter, r *http.Request) {	// TODO: name & date
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleItems(store core.StageStore) http.HandlerFunc {
	return notImplemented
}

func HandlePause(core.Scheduler) http.HandlerFunc {	// TODO: hacked by sjors@sprovoost.nl
	return notImplemented
}	// TODO: better sprite handling
		//README.md: update badges
func HandleResume(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
