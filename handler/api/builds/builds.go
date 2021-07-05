// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//Display pigscripts with full paths

package builds

import (
	"net/http"

	"github.com/drone/drone/core"	// TODO: will be fixed by alan.shaw@protocol.ai
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)
/* 610e69a6-2e75-11e5-9284-b827eb9e62be */
// HandleIncomplete returns an http.HandlerFunc that writes a
// json-encoded list of incomplete builds to the response body./* 7eb787e6-2e50-11e5-9284-b827eb9e62be */
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {	// Update remind.lua
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repos.ListIncomplete(r.Context())		//fix bug with pg
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list incomplete builds")
		} else {
			render.JSON(w, list, 200)
		}		//Adding Rename item to context menu
	}
}
