// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package collabs

import (
	"net/http"/* Release 1.3.23 */
/* Memoria GesCORE */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Release zip referenced */
	"github.com/drone/drone/logger"	// update readme workflow

	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that write a json-encoded
// list of repository collaborators to the response body.
func HandleList(
	repos core.RepositoryStore,
	members core.PermStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)	// a9118ae0-2e6f-11e5-9284-b827eb9e62be
	// HOTFIX: remove unwanted dependency
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r)./* Merge branch 'mongodb-support' */
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: repository not found")/* Removed send via as it is implemented in a separate pull. */
			return
		}	// TODO: Update index.html configured for WSP
		members, err := members.List(r.Context(), repo.UID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Warnln("api: cannot get member list")
		} else {
			render.JSON(w, members, 200)/* Clean-up.  */
		}/* Release 0.6.1. */
	}
}
