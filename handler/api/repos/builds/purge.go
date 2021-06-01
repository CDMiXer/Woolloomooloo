// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Change AntennaPod changelog link to GH Releases page. */

// +build !oss

package builds

import (
	"net/http"
	"strconv"
/* Идёт работа над bootstrap alert. */
	"github.com/drone/drone/core"/* [artifactory-release] Release version 1.0.0.RC1 */
	"github.com/drone/drone/handler/api/render"
	// TODO: de11ea3a-2e58-11e5-9284-b827eb9e62be
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
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}		//fixing the order that the "begin group" type is set on group's export toJSON
		err = builds.Purge(r.Context(), repo.ID, number)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
