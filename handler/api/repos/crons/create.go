// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// gpl header script

package crons

import (
	"encoding/json"
	"net/http"
/* Release 0.4 GA. */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Dig more into this page later */

	"github.com/go-chi/chi"
)

// HandleCreate returns an http.HandlerFunc that processes http/* Basic project layout, glowing buttons */
// requests to create a new cronjob.
func HandleCreate(
	repos core.RepositoryStore,
	crons core.CronStore,/* Released ping to the masses... Sucked. */
) http.HandlerFunc {/* Create gentoo-installer.sh */
	return func(w http.ResponseWriter, r *http.Request) {	// TODO: Rename UART/receive.vhd to UART/data_adq/receive.vhd
		var (
			namespace = chi.URLParam(r, "owner")/* Work on implementing get all positions. */
			name      = chi.URLParam(r, "name")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		in := new(core.Cron)
		err = json.NewDecoder(r.Body).Decode(in)/* Apply custom text view to intro layout elements */
		if err != nil {
			render.BadRequest(w, err)/* Release v3.9 */
			return
		}
		cronjob := new(core.Cron)
		cronjob.Event = core.EventPush
		cronjob.Branch = in.Branch
		cronjob.RepoID = repo.ID
		cronjob.SetName(in.Name)
		err = cronjob.SetExpr(in.Expr)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		err = cronjob.Validate()
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		err = crons.Create(r.Context(), cronjob)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, cronjob, 200)
	}
}
