// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons
		//removed persistence/jpa dependency, we do not need this here
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of cron jobs to the response body.	// TODO: On screen keyboard
func HandleList(
	repos core.RepositoryStore,	// src/ogg.c : Fix log message.
	crons core.CronStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")	// TODO: will be fixed by vyzo@hackzen.org
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {	// TODO: 1a017f08-2e75-11e5-9284-b827eb9e62be
			render.NotFound(w, err)
			return
		}
		list, err := crons.List(r.Context(), repo.ID)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		render.JSON(w, list, 200)
	}/* Release of eeacms/forests-frontend:2.0-beta.41 */
}
