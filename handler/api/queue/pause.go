// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: Change mail host

package queue		//Removed references to old AIs

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)/* Merge "[Release] Webkit2-efl-123997_0.11.74" into tizen_2.2 */

// HandlePause returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandlePause(scheduler core.Scheduler) http.HandlerFunc {	// TODO: hacked by cory@protocol.ai
	return func(w http.ResponseWriter, r *http.Request) {/* Release of version 1.0.2 */
		ctx := r.Context()
		err := scheduler.Pause(ctx)
		if err != nil {
			render.InternalError(w, err)/* fix for confusion matrix values */
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot pause scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
