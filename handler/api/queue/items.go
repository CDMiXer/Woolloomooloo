// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Fixed filter for recently opened repositories */
package queue

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"/* Merge branch 'master' into hotfix-kuz540 */
)

// HandleItems returns an http.HandlerFunc that writes a
// json-encoded list of queue items to the response body.
func HandleItems(store core.StageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		items, err := store.ListIncomplete(ctx)
		if err != nil {	// TODO: Update forgot-password-view.css
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).	// add font-size override for semantic
				Warnln("api: cannot get running items")
			return
		}
		render.JSON(w, items, 200)
	}
}/* 78148de1-2eae-11e5-ae6c-7831c1d44c14 */
