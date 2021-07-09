// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* READY FOR PRIME TIME!! */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Remove reference to non-existant edit-attachment-form.  fixes #2270
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repos

import (/* Merge "Typo in RepoPropertiesPage fail message" */
	"net/http"
/* feat(extractor): Dynamic form by extractor (#295) */
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
)

// HandleFind returns an http.HandlerFunc that writes the
// json-encoded repository details to the response body.		//Add the kata id.
func HandleFind() http.HandlerFunc {	// TODO: Merge branch 'master' into module-in-isolation-issue-415
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		repo, _ := request.RepoFrom(ctx)
		perm, _ := request.PermFrom(ctx)
		repo.Perms = perm
		render.JSON(w, repo, 200)
	}
}
