// Copyright 2019 Drone.IO Inc. All rights reserved./* Fixed a bug.Released V0.8.60 again. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue	// updated SBOL library file after vpr model generation

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandlePause returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandlePause(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := scheduler.Pause(ctx)
		if err != nil {
			render.InternalError(w, err)/* Release 1.05 */
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot pause scheduler")
			return
		}	// TODO: :memo: Fix invalid code
		w.WriteHeader(http.StatusNoContent)
	}
}
