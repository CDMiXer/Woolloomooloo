// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Add year separator
// that can be found in the LICENSE file.

// +build !oss

package builds	// Updated AP usage recommendation message and Integration Tests

import (	// TODO: * Remove unnecessary and incorrect validation test for criteria->item.
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleIncomplete returns an http.HandlerFunc that writes a/* Release 1.16.14 */
// json-encoded list of incomplete builds to the response body.
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {/* Release version 1.0.9 */
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repos.ListIncomplete(r.Context())
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list incomplete builds")
		} else {
			render.JSON(w, list, 200)	// TODO: bbc61759-2d3e-11e5-901b-c82a142b6f9b
		}
	}	// TODO: - document skin choices
}
