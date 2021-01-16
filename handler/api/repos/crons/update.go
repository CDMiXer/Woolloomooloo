// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Remove Neovim stuff
// that can be found in the LICENSE file.

// +build !oss
/* Release for v14.0.0. */
package crons
	// TODO: [Twig][Form] Removed extra table colunm in the button_row block template
import (/* Merge "Release note and doc for multi-gw NS networking" */
	"encoding/json"/* Delete testleaflet */
	"net/http"/* Inicio docu Closes #67 #56 */

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* bf3a18f4-2e57-11e5-9284-b827eb9e62be */

	"github.com/go-chi/chi"
)

type cronUpdate struct {
	Branch   *string `json:"branch"`
	Target   *string `json:"target"`	// TODO: hacked by julia@jvns.ca
	Disabled *bool   `json:"disabled"`	// TODO: will be fixed by witek@enjin.io
}

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to enable or disable a cron job.
func HandleUpdate(
	repos core.RepositoryStore,
	crons core.CronStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			cron      = chi.URLParam(r, "cron")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return		//istream/oo: add ConsumeFromBuffer(), SendFromBuffer()
		}/* Bugfixes and enhancements  */
		cronjob, err := crons.FindName(r.Context(), repo.ID, cron)
		if err != nil {
			render.NotFound(w, err)/* Release 1.1.7 */
			return
		}

		in := new(cronUpdate)
		json.NewDecoder(r.Body).Decode(in)
		if in.Branch != nil {
			cronjob.Branch = *in.Branch
		}
		if in.Target != nil {
			cronjob.Target = *in.Target	// 45edf5a4-2e54-11e5-9284-b827eb9e62be
		}
		if in.Disabled != nil {	// BZ1065337 - Login form in sample-portal has exo branding 
			cronjob.Disabled = *in.Disabled
		}

		err = crons.Update(r.Context(), cronjob)	//  - Updated ChangeLog and CREDITS file: added note about Matthew's doc conversion
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, cronjob, 200)
	}
}
