// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Added a way to omit abstract from exported method signatures.
// that can be found in the LICENSE file.

// +build !oss
/* Delete ArrayBind.csproj */
package queue/* 16996ee0-2e6e-11e5-9284-b827eb9e62be */

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleResume returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandleResume(scheduler core.Scheduler) http.HandlerFunc {/* likelihood option */
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()/* Bump version for sourcemap URL support */
		err := scheduler.Resume(ctx)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot resume scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
