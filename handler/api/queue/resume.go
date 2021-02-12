// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//#25 Android scroll issue

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
		ctx := r.Context()/* Release 0.6.7. */
		err := scheduler.Resume(ctx)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot resume scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)/* Merge "ARM: dts: msm: Add smb_stat pinctrl node for mdmcalifornium" */
	}/* Release 1.6.15 */
}
