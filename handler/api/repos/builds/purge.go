// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package builds/* Release 3.5.0 */

import (
	"net/http"
	"strconv"	// Show Legend inside pathway diagram (fixes #979)
/* for isEmptyToNullReading */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"/* Create DEPRECATED -Ubuntu Gnome Rolling Release */
)/* Release version 0.4 */
	// postfix: fix test config dir
// HandlePurge returns an http.HandlerFunc that purges the
// build history. If successful a 204 status code is returned.
func HandlePurge(repos core.RepositoryStore, builds core.BuildStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			before    = r.FormValue("before")
		)
		number, err := strconv.ParseInt(before, 10, 64)/* Same optimization level for Debug & Release */
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		err = builds.Purge(r.Context(), repo.ID, number)
		if err != nil {
			render.InternalError(w, err)/* Update Releases */
			return
		}/* Release 0.3.0. Add ip whitelist based on CIDR. */
		w.WriteHeader(http.StatusNoContent)
	}
}
