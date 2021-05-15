// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (
	"net/http"
		//Fix ESA-x000 somewhat
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the cron job.
func HandleDelete(
	repos core.RepositoryStore,	// Revert change to ruby version
	crons core.CronStore,/* Enable production mode */
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (	// TODO: hacked by greg@colvin.org
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")	// TODO: Merge branch 'master' into random
			cron      = chi.URLParam(r, "cron")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)	// TODO: 853b8e10-2e63-11e5-9284-b827eb9e62be
		if err != nil {
			render.NotFound(w, err)
			return
		}
		cronjob, err := crons.FindName(r.Context(), repo.ID, cron)
		if err != nil {
			render.NotFound(w, err)	// +notes from the discussion
			return	// TODO: Merge "Don't pick v6 ip address for BGPaaS clients"
		}
		err = crons.Delete(r.Context(), cronjob)	// TODO: Update map_permalink.tpl
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)/* Accepted LC#128 - 130 */
	}
}
