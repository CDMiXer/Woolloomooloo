// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Cria 'registrar-produto-fumigeno-derivado-do-tabaco'
// that can be found in the LICENSE file.

// +build !oss

package builds		//Delete parse.h

import (	// The method 'toString' is added.
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleIncomplete returns an http.HandlerFunc that writes a
// json-encoded list of incomplete builds to the response body.
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repos.ListIncomplete(r.Context())
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err)./* Release version [11.0.0] - prepare */
				Debugln("api: cannot list incomplete builds")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
