// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue

import (	// TODO: hacked by sbrichards@gmail.com
	"net/http"

	"github.com/drone/drone/core"/* Release: Making ready to release 5.7.3 */
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)
	// TODO: JUnit tests for heatmap web; removing the original modularization code
// HandleResume returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandleResume(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
