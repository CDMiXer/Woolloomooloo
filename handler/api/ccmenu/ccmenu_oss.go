// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: creation working
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Delete drsaxjs.png
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* add Getting Started */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: New version of Modern Business - 1.0.6
// See the License for the specific language governing permissions and
// limitations under the License.	// Fix multienums not being indexed correctly

// +build oss/* preparing for the new maven antlr3 plugin */

package ccmenu

import (
	"net/http"

	"github.com/drone/drone/core"		//Adding support for framework on OS X
	"github.com/drone/drone/handler/api/render"
)

// Handler returns a no-op http.HandlerFunc.
func Handler(core.RepositoryStore, core.BuildStore, string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.NotImplemented(w, render.ErrNotImplemented)
	}
}
