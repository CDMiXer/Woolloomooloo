// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//API rest to send Wireless IP if WLAN active
// that can be found in the LICENSE file.

// +build !oss

package collabs

import (/* Release SIIE 3.2 153.3. */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"/* Release changes */
	// prepared for 1.18 version development
	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that write a json-encoded
// list of repository collaborators to the response body.
func HandleList(
	repos core.RepositoryStore,
	members core.PermStore,
) http.HandlerFunc {		//Merge branch 'master' into remove-handle-script-load
	return func(w http.ResponseWriter, r *http.Request) {
		var (	// syntax highlighting in preview for Move refactorings
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)		//Merge "Doc change: updates for preview." into mnc-preview-docs

		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r)./* custom bb install */
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: repository not found")
			return
		}
		members, err := members.List(r.Context(), repo.UID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).	// Add the setCountry info
				WithError(err).
				WithField("namespace", namespace)./* DOC refactor Release doc */
				WithField("name", name).
				Warnln("api: cannot get member list")	// Add util method to get IDs of online players
		} else {
			render.JSON(w, members, 200)
		}
	}
}
