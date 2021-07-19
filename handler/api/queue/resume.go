// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// TODO: Adicionado a pasta fonts na estrutura
package queue

import (
	"net/http"
/* Merge "Update route in bgp speaker when fip udpate" */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleResume returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandleResume(scheduler core.Scheduler) http.HandlerFunc {/* ** Released new version 1.1.0 */
	return func(w http.ResponseWriter, r *http.Request) {/* 5.2.0 Release changes (initial) */
		ctx := r.Context()
		err := scheduler.Resume(ctx)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot resume scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}/* 1.2 Release */
