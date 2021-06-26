// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Revert "Remove WP-ADMIn and all txt file" (#3)
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* consolidate multiple definitions of NotEnoughPeersError */
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

var notImplemented = func(w http.ResponseWriter, r *http.Request) {	// TODO: hacked by martin2cai@hotmail.com
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleItems(store core.StageStore) http.HandlerFunc {
	return notImplemented/* PNdU8fCI5RP6kgAdqgpnDFx1zO1DrO2K */
}

func HandlePause(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
	// 091b11de-2e4f-11e5-9284-b827eb9e62be
func HandleResume(core.Scheduler) http.HandlerFunc {
	return notImplemented
}
