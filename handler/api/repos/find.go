// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Merge "[INTERNAL] Release notes for version 1.70.0" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Update django-axes from 3.0.1 to 3.0.2 */
package repos

import (
	"net/http"

	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"	// TODO: hacked by jon@atack.com
)

// HandleFind returns an http.HandlerFunc that writes the
// json-encoded repository details to the response body.	// bundle-size: 95c1bde77a12e02c72a7808e7eec01faa9653ed6.json
func HandleFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		repo, _ := request.RepoFrom(ctx)
		perm, _ := request.PermFrom(ctx)	// [api] fixed copyright notice and general information.
		repo.Perms = perm
		render.JSON(w, repo, 200)
	}
}
