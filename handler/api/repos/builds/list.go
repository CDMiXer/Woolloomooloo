// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//remove -fvia-C that I apparrently accidentally added recently
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Added third parallel version and customized output file names
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Fixed missing selectedValue */
/* Release new version 2.2.6: Memory and speed improvements (famlam) */
package builds

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
		//[IMP] document : added missing filter string in search view.
	"github.com/go-chi/chi"	// TODO: hacked by vyzo@hackzen.org
)
/* Add headers method to set multiple headers at once */
// HandleList returns an http.HandlerFunc that writes a json-encoded	// TODO: hacked by ng8eke@163.com
// list of build history to the response body.
func HandleList(
	repos core.RepositoryStore,
	builds core.BuildStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			branch    = r.FormValue("branch")
			page      = r.FormValue("page")
			perPage   = r.FormValue("per_page")
		)
		offset, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(perPage)
		if limit < 1 || limit > 100 {
			limit = 25
		}/* For #943 Removed ${argLine}, it was crashing VM when running testes. */
		switch offset {
		case 0, 1:	// Commented out a compilation error
			offset = 0/* (v2) Chains: show chains and examples in the same table. */
		default:
			offset = (offset - 1) * limit
		}
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {	// TODO: Implement first pass at building CSS and HTML
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).		//groovy-all.jar may require some other jars to run groovy compiler
				WithField("name", name).		//doc: update image
				Debugln("api: cannot find repository")
			return
		}
/* - git ignore for crawler */
		var results []*core.Build		//Add thickbox to pages that use media uploader.
		if branch != "" {
			ref := fmt.Sprintf("refs/heads/%s", branch)
			results, err = builds.ListRef(r.Context(), repo.ID, ref, limit, offset)
		} else {
			results, err = builds.List(r.Context(), repo.ID, limit, offset)
		}

		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot list builds")
		} else {
			render.JSON(w, results, 200)
		}
	}
}
