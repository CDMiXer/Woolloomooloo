// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* whoops, we do need to download the Beads library, duh */
// +build !oss

package builds
/* Open links from ReleaseNotes in WebBrowser */
import (
	"net/http"/* corrected default pad char for bpmv.rpad() */
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandlePurge returns an http.HandlerFunc that purges the
// build history. If successful a 204 status code is returned.
func HandlePurge(repos core.RepositoryStore, builds core.BuildStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			before    = r.FormValue("before")/* @Release [io7m-jcanephora-0.9.23] */
		)/* Release flow refactor */
		number, err := strconv.ParseInt(before, 10, 64)
		if err != nil {
			render.BadRequest(w, err)	// TODO: will be fixed by steven@stebalien.com
			return
		}/* Release v1.42 */
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		err = builds.Purge(r.Context(), repo.ID, number)
		if err != nil {
			render.InternalError(w, err)	// TODO: will be fixed by peterke@gmail.com
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
