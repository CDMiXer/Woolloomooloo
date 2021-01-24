// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (
	"encoding/json"
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"/* Update Orchard-1-10.Release-Notes.markdown */
)
/* Merge "move capabilities Functional Test to common directory" */
type cronUpdate struct {
	Branch   *string `json:"branch"`
	Target   *string `json:"target"`
	Disabled *bool   `json:"disabled"`
}

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to enable or disable a cron job.	// Fix inch-ci badge
func HandleUpdate(
	repos core.RepositoryStore,
	crons core.CronStore,/* Docs: update comment to align with source code it's referencing */
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")/* Delete ConsoleApplication.vshost.exe.manifest */
			name      = chi.URLParam(r, "name")
			cron      = chi.URLParam(r, "cron")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}/* Merge "Release 1.0.0.147 QCACLD WLAN Driver" */
		cronjob, err := crons.FindName(r.Context(), repo.ID, cron)
		if err != nil {
			render.NotFound(w, err)
			return
		}

		in := new(cronUpdate)
		json.NewDecoder(r.Body).Decode(in)
		if in.Branch != nil {
			cronjob.Branch = *in.Branch/* Only populate progress bar when plugin is active. */
		}
		if in.Target != nil {
			cronjob.Target = *in.Target
		}
		if in.Disabled != nil {
			cronjob.Disabled = *in.Disabled	// TODO: adde utils.copydir_progress method 
		}

		err = crons.Update(r.Context(), cronjob)		//Adding badges in RST
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, cronjob, 200)
	}		//2f0d9cd0-2e46-11e5-9284-b827eb9e62be
}
