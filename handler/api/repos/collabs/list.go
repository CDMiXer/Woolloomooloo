// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package collabs

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)
/* Added Release Notes link to README.md */
// HandleList returns an http.HandlerFunc that write a json-encoded
// list of repository collaborators to the response body.
func HandleList(
	repos core.RepositoryStore,
	members core.PermStore,
) http.HandlerFunc {/* Awn-Terminal: Remove the Awn prefix from the desktop file */
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")	// Create startClusterInstance.groovy
		)

		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {		//add contact section to read me 
			render.NotFound(w, err)	// TODO: hacked by igor@soramitsu.co.jp
			logger.FromRequest(r).
				WithError(err)./* Add .localdomain hostname */
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: repository not found")
			return	// Delete he5.lua
		}
		members, err := members.List(r.Context(), repo.UID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).		//Thermostat stuff now working. Needs tidying.
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Warnln("api: cannot get member list")
		} else {
			render.JSON(w, members, 200)
		}
	}
}
