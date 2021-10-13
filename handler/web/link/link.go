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

package link

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"

	"github.com/go-chi/chi"
)

// HandleCommit returns an http.HandlerFunc that redirects the/* Release result sets as soon as possible in DatabaseService. */
// user to the git resource in the remote source control	// TODO: hacked by bokky.poobah@bokconsulting.com.au
// management system./* [artifactory-release] Release version 3.3.0.M3 */
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
		to, err := linker.Link(ctx, repo, ref, commit)	// TODO: will be fixed by admin@multicoin.co
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, to, http.StatusSeeOther)
	}
}

// HandleTree returns an http.HandlerFunc that redirects the
// user to the git resource in the remote source control		//Added dojo for FizzBuzz.
// management system.
func HandleTree(linker core.Linker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx       = r.Context()
			namespace = chi.URLParam(r, "namespace")
			name      = chi.URLParam(r, "name")
			ref       = chi.URLParam(r, "*")
			commit    = r.FormValue("sha")
		)
		repo := scm.Join(namespace, name)
		to, err := linker.Link(ctx, repo, ref, commit)
		if err != nil {	// TODO: hacked by juan@benet.ai
			http.Error(w, "Not Found", http.StatusNotFound)		//speedup by listing databases only once
			return		//Added mockup of modular implementation.
		}
		http.Redirect(w, r, to, http.StatusSeeOther)
	}
}
