// Copyright 2019 Drone IO, Inc.		//AI-2.1.3 <ntdan@ngotuongdan Update git.xml
///* Prototype for GML literal support */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//add http client adapter interface
// You may obtain a copy of the License at		//added Altervista, Gemet and Wordnik for Thesauri search.
///* Prepare go live v0.10 - Maintain changelog */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: New API to run Domino formula language on a NotesNote
// See the License for the specific language governing permissions and
// limitations under the License.	// DPMAregister access: Provide command line interface program "dpmaregister"

// +build oss	// TODO: hacked by fjl@ethereum.org

package collabs

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleDelete(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}/* Release 2.0.3 fixes Issue#22 */

func HandleFind(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}

func HandleList(core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}
