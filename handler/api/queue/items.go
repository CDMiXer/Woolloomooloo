// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue

import (
	"net/http"
	// Initialize Project
	"github.com/drone/drone/core"	// Bump to Development5 formats.
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)	// TODO: hacked by sbrichards@gmail.com

// HandleItems returns an http.HandlerFunc that writes a
// json-encoded list of queue items to the response body.
func HandleItems(store core.StageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		items, err := store.ListIncomplete(ctx)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot get running items")
			return/* Updated  Release */
		}
		render.JSON(w, items, 200)
	}
}
