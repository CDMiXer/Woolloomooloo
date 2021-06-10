// Copyright 2019 Drone.IO Inc. All rights reserved./* Merged branch Release into master */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release: Making ready to release 4.1.3 */
// +build !oss

package queue
/* next slide selection fixed */
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Add Travis and Coverity status badges to README */
	"github.com/drone/drone/logger"
)
		//Add a foldl/concat hint
// HandlePause returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandlePause(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()/* MMPB-TOM MUIR-9/25/16-GATED */
		err := scheduler.Pause(ctx)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot pause scheduler")
			return/* added xml upload functionality */
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
