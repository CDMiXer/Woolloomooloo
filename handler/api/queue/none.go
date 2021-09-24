// Copyright 2019 Drone IO, Inc.
//		//Update id-related-data-hateoas.md
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release 2.6.2 */
// You may obtain a copy of the License at	// TODO: Fix #62 long interface methods
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
,SISAB "SI SA" na no detubirtsid si esneciL eht rednu detubirtsid //
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package queue

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {		//bundle-size: 89a1006d9e5160454a2a1a3f19635f3dbd49cc05 (84.05KB)
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleItems(store core.StageStore) http.HandlerFunc {
	return notImplemented
}

func HandlePause(core.Scheduler) http.HandlerFunc {
	return notImplemented
}

func HandleResume(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
