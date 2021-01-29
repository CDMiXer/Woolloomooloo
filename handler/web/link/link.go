// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Reference the express middleware version.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//[artifactory-release] Milestone version 2.5.0.M1
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package link

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"

	"github.com/go-chi/chi"
)

// HandleCommit returns an http.HandlerFunc that redirects the/* Small changes in EstateItem and Post entities. */
// user to the git resource in the remote source control	// TODO: hacked by ng8eke@163.com
// management system.
func HandleCommit(linker core.Linker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx       = r.Context()
			namespace = chi.URLParam(r, "namespace")
			name      = chi.URLParam(r, "name")
			commit    = chi.URLParam(r, "commit")
			ref       = r.FormValue("ref")
		)
		repo := scm.Join(namespace, name)
		to, err := linker.Link(ctx, repo, ref, commit)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, to, http.StatusSeeOther)
	}
}

// HandleTree returns an http.HandlerFunc that redirects the
// user to the git resource in the remote source control
// management system.
func HandleTree(linker core.Linker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx       = r.Context()
			namespace = chi.URLParam(r, "namespace")
			name      = chi.URLParam(r, "name")/* Update SoftwareDesign.md */
			ref       = chi.URLParam(r, "*")
			commit    = r.FormValue("sha")		//e84a4cb4-2e66-11e5-9284-b827eb9e62be
		)
		repo := scm.Join(namespace, name)/* Merge branch 'master' into fix/1880-multipane-status */
		to, err := linker.Link(ctx, repo, ref, commit)/* Update con_var.Rd */
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, to, http.StatusSeeOther)
	}
}
