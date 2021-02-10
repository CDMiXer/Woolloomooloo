// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Fix unit test after changing EquipmentMessageSender class */

// +build !oss

package crons
	// fix prod secure url
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	// TODO: Merge branch 'master' into racemodeUI
	"github.com/go-chi/chi"
)/* Release 2.0.0 of PPWCode.Util.AppConfigTemplate */

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the cron job./* Release notes for 1.0.91 */
func HandleDelete(
	repos core.RepositoryStore,
	crons core.CronStore,
) http.HandlerFunc {/* added sane application path */
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")/* Release 1-110. */
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
		err = crons.Delete(r.Context(), cronjob)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
