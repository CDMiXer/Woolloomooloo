// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* lost in merge */

sso! dliub+ //

package crons

import (	// TODO: hacked by brosner@gmail.com
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of cron jobs to the response body.
func HandleList(
	repos core.RepositoryStore,	// for issue #5
	crons core.CronStore,		//Converted all Entity functions to use FP class math
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)/* binary Release */
			return
		}
		list, err := crons.List(r.Context(), repo.ID)
		if err != nil {
			render.NotFound(w, err)
			return
		}/* 10c7104c-2e57-11e5-9284-b827eb9e62be */
		render.JSON(w, list, 200)
	}
}
