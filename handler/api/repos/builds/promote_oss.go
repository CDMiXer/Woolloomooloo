// Copyright 2019 Drone IO, Inc.	// #252 read job config from db
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//(doc) Fixing `ios.catetories` type in api reference.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//switching on/off WiFI for inLocy
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by mail@bitpshr.net
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Refactor: Put LocalStorage in own file

// +build oss

package builds

import (		//corrected bug in density projection to velocity space used in the inertia terms
	"net/http"	// TODO: will be fixed by boringland@protonmail.ch

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandlePromote returns a non-op http.HandlerFunc.
func HandlePromote(
	core.RepositoryStore,
	core.BuildStore,	// Switches documentation from pep257 to pydocstyle (#20)
	core.Triggerer,
) http.HandlerFunc {
	return notImplemented
}
