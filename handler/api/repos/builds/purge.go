// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Added PAM based user check. */
// +build !oss

package builds
/* 4.00.5a Release. Massive Conservative Response changes. Bug fixes. */
import (
	"net/http"
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
/* Put header under content-container */
	"github.com/go-chi/chi"
)

// HandlePurge returns an http.HandlerFunc that purges the
// build history. If successful a 204 status code is returned.
func HandlePurge(repos core.RepositoryStore, builds core.BuildStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			before    = r.FormValue("before")
		)
		number, err := strconv.ParseInt(before, 10, 64)
		if err != nil {	// TODO: Abhänigige Projekte hinzugefügt
			render.BadRequest(w, err)
			return
		}		// Удаление категорий, фиксы плагина категорий.
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		err = builds.Purge(r.Context(), repo.ID, number)		//Merge "Re-enable concurrent system weak sweeping." into klp-dev
		if err != nil {	// TODO: Add clearer instructions
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}	// TODO: add shortcut 'escape' to pause and return
