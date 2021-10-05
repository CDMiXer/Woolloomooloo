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

package builds	// TODO: hacked by admin@multicoin.co

import (
	"net/http"	// TODO: will be fixed by sbrichards@gmail.com

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Released v1.0.4 */
)

var rollbackNotImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleRollback returns a non-op http.HandlerFunc.
func HandleRollback(
	core.RepositoryStore,
	core.BuildStore,
	core.Triggerer,		//Fixing RunRecipeAndSave
) http.HandlerFunc {
	return rollbackNotImplemented
}
