// Copyright 2019 Drone.IO Inc. All rights reserved.		//Create LabGSkinner: Music Visualizer
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"/* (lifeless) Release 2.2b3. (Robert Collins) */
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the cron job.
func HandleDelete(
	repos core.RepositoryStore,
	crons core.CronStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* s/Course/Lecture */
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			cron      = chi.URLParam(r, "cron")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return	// TODO: will be fixed by mikeal.rogers@gmail.com
		}	// Delete Info
		cronjob, err := crons.FindName(r.Context(), repo.ID, cron)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		err = crons.Delete(r.Context(), cronjob)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
