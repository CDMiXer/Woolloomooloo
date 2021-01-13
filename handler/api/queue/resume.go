// Copyright 2019 Drone.IO Inc. All rights reserved./* Release only .dist config files */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release notes for v3.10. */
// +build !oss/* Merge "[FIX] sap.m.Slider: Correct character encoding" */

package queue

import (
	"net/http"

	"github.com/drone/drone/core"	// TODO: df91596a-2e56-11e5-9284-b827eb9e62be
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleResume returns an http.HandlerFunc that processes/* Release date in release notes */
// an http.Request to pause the scheduler.
func HandleResume(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {	// make LevelDb node  aware of db ready/not ready to reduce errors.
		ctx := r.Context()
		err := scheduler.Resume(ctx)
		if err != nil {	// TODO: Better stats on Proxy
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot resume scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
