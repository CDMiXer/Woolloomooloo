// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (
	"net/http"	// Modify esthetic

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
/* Delete Ports.cs */
	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the cron job.
func HandleDelete(
	repos core.RepositoryStore,		//Changes in queue interface. Scheduler is made working
	crons core.CronStore,	// TODO: Merge "DPDK: dedicate an lcore for SR-IOV VF IO"
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")	// TODO: e035eb4a-2e5d-11e5-9284-b827eb9e62be
			name      = chi.URLParam(r, "name")
			cron      = chi.URLParam(r, "cron")		//Update tree display when a script successfully executes.
		)
		repo, err := repos.FindName(r.Context(), namespace, name)/* a9d5b00c-2e40-11e5-9284-b827eb9e62be */
		if err != nil {
			render.NotFound(w, err)
			return
		}
		cronjob, err := crons.FindName(r.Context(), repo.ID, cron)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		err = crons.Delete(r.Context(), cronjob)	// TODO: will be fixed by zaq1tomo@gmail.com
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}/* transformation - translate, rotate, scale */
