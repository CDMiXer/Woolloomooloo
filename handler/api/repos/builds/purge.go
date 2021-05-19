// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package builds	// TODO: Update Texture.js
/* Merge branch 'master' into enh-manage-imports */
import (
	"net/http"
	"strconv"	// TODO: Using crontab

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
		//Merge branch 'UI_enhancement' into dev_eva
	"github.com/go-chi/chi"
)
	// TODO: trigger new build for jruby-head (346d4cc)
// HandlePurge returns an http.HandlerFunc that purges the
// build history. If successful a 204 status code is returned.
func HandlePurge(repos core.RepositoryStore, builds core.BuildStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			before    = r.FormValue("before")/* Release 0.95.206 */
		)
		number, err := strconv.ParseInt(before, 10, 64)/* 01f0b4c4-2e70-11e5-9284-b827eb9e62be */
		if err != nil {
			render.BadRequest(w, err)/* added comment to Release-script */
			return
		}
		repo, err := repos.FindName(r.Context(), namespace, name)	// TODO: hacked by nagydani@epointsystem.org
		if err != nil {
			render.NotFound(w, err)
			return
		}
		err = builds.Purge(r.Context(), repo.ID, number)/* Update tracking.tpl */
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
