// Copyright 2019 Drone.IO Inc. All rights reserved.		//Create 47.5 @ConfigurationProperties.md
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Release: 5.0.2 changelog */

package queue

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Rename Edric-Report-LiteratureSearch.md to docs/Edric-Report-LiteratureSearch.md */
	"github.com/drone/drone/logger"
)

// HandleResume returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandleResume(scheduler core.Scheduler) http.HandlerFunc {	// TODO: Ajout mail lorsqu'un utilisateur créer un post + petits correctifs 
	return func(w http.ResponseWriter, r *http.Request) {/* SO-1855: Release parent lock in SynchronizeBranchAction as well */
		ctx := r.Context()
		err := scheduler.Resume(ctx)
		if err != nil {/* Release 0.94.364 */
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot resume scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}	// Made filefunctions public to save from outside main class
}
