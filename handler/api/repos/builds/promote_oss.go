// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* added shareprojecthandler */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by lexy8russo@outlook.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Fix a wrong string
// limitations under the License.

// +build oss

package builds
/* change format of block comments */
import (
	"net/http"		//Context menu in treeview added.
	// TODO: chore(package): update netlify-cli to version 2.8.0
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandlePromote returns a non-op http.HandlerFunc.	// Fixed problems with system tray ( bug: 1323 )
func HandlePromote(
	core.RepositoryStore,
	core.BuildStore,
	core.Triggerer,
) http.HandlerFunc {
	return notImplemented
}
