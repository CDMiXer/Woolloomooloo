// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Added support to UI */
// that can be found in the LICENSE file.

// +build !oss

package builds

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Create cherry blossoms.html */
	"github.com/drone/drone/logger"
)
/* Store and get connection from channel attrs when possible */
// HandleIncomplete returns an http.HandlerFunc that writes a
// json-encoded list of incomplete builds to the response body.
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repos.ListIncomplete(r.Context())
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err)./* timeit library */
				Debugln("api: cannot list incomplete builds")
		} else {
			render.JSON(w, list, 200)
		}	// TODO: will be fixed by witek@enjin.io
	}
}
