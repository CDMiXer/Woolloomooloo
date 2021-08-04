// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: Better cloning of the original callstack

// +build !oss

package queue

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandlePause returns an http.HandlerFunc that processes	// TODO: Removed README colored alerts section
// an http.Request to pause the scheduler.	// TODO: will be fixed by zaq1tomo@gmail.com
func HandlePause(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := scheduler.Pause(ctx)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot pause scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)		//chore(deps): update dependency lint-staged to ^8.1.4
	}
}
