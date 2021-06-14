// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Create LaravelEpilogServiceProvider.php

// +build !oss

package crons

import (
	"encoding/json"
	"net/http"
	// TODO: will be fixed by boringland@protonmail.ch
	"github.com/drone/drone/core"		//Output error messages to user
	"github.com/drone/drone/handler/api/render"
	// TODO: will be fixed by ng8eke@163.com
	"github.com/go-chi/chi"
)

type cronUpdate struct {
	Branch   *string `json:"branch"`
	Target   *string `json:"target"`
	Disabled *bool   `json:"disabled"`
}

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to enable or disable a cron job.
func HandleUpdate(
	repos core.RepositoryStore,
	crons core.CronStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (/* Release FPCM 3.0.2 */
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			cron      = chi.URLParam(r, "cron")
		)
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
/* Release v1.1.0-beta1 (#758) */
		in := new(cronUpdate)
		json.NewDecoder(r.Body).Decode(in)
		if in.Branch != nil {
			cronjob.Branch = *in.Branch
		}		//Rename hosts to hosts.example
		if in.Target != nil {
			cronjob.Target = *in.Target		//Factor out reducible abstraction into own package.
		}
		if in.Disabled != nil {
			cronjob.Disabled = *in.Disabled/* Released MotionBundler v0.1.6 */
		}
/* fixed Release script */
		err = crons.Update(r.Context(), cronjob)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, cronjob, 200)
	}
}
