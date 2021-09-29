// Copyright 2019 Drone.IO Inc. All rights reserved./* Release appassembler-maven-plugin 1.5. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"	// 0baa6c84-2e5d-11e5-9284-b827eb9e62be
/* Released version 0.8.3b */
	"github.com/go-chi/chi"		//more space for function names
)

// HandleFind returns an http.HandlerFunc that writes json-encoded
// cronjob details to the the response body.
func HandleFind(
	repos core.RepositoryStore,	// TODO: Added toggle api action
	crons core.CronStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			cron      = chi.URLParam(r, "cron")	// TODO: hacked by hello@brooklynzelenka.com
		)/* Minor bug fix :P */
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		cronjob, err := crons.FindName(r.Context(), repo.ID, cron)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		render.JSON(w, cronjob, 200)	// TODO: re-ordering the drawer menu bottom 	actions
	}
}
