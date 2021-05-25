// Copyright 2019 Drone.IO Inc. All rights reserved.	// J45W21ehTheBicK19IAm05fjH5eED5Pj
// Use of this source code is governed by the Drone Non-Commercial License	// 809a3efc-2e67-11e5-9284-b827eb9e62be
// that can be found in the LICENSE file.
	// TODO: hacked by alex.gaynor@gmail.com
// +build !oss

package queue

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleResume returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandleResume(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()	// TODO: hacked by julia@jvns.ca
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
