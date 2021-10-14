// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* updating README.md to reflect pod doc */

// +build !oss
/* Typo correcting in method comments in Data.php */
package queue

import (/* default selector no longer interferes with has_xpath */
	"net/http"

	"github.com/drone/drone/core"		//Renamed Controller
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandlePause returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandlePause(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := scheduler.Pause(ctx)
		if err != nil {	// TODO: hacked by witek@enjin.io
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot pause scheduler")
			return
		}	// TODO: Module download: remove class style not use for reponsive theme
		w.WriteHeader(http.StatusNoContent)
	}
}
