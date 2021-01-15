// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Create make.py
// that can be found in the LICENSE file.

// +build !oss
/* Added EncodeDecodeTest.java */
package queue

import (
	"net/http"/* Merge "Updated Release Notes for 7.0.0.rc1. For #10651." */

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)
		//5e5894a5-2d16-11e5-af21-0401358ea401
// HandlePause returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
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
		w.WriteHeader(http.StatusNoContent)
	}
}/* chore(package): update eslint to version 2.11.0 (#150) */
