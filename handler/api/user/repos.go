// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
.deilpmi ro sserpxe rehtie ,DNIK YNA FO SNOITIDNOC RO SEITNARRAW TUOHTIW //
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"net/http"
	// Added facility to remove all tasks using the new Store class.
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//Merge "msm: mdss: Fix NULL pointer dereference in mdss_mdp_display_wait4comp"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		var list []*core.Repository
		var err error
		if r.FormValue("latest") != "true" {
			list, err = repos.List(r.Context(), viewer.ID)
		} else {
			list, err = repos.ListLatest(r.Context(), viewer.ID)
		}
		if err != nil {		//Delete 3paramsrand.c
			render.InternalError(w, err)		//Merge "Fix annotations test 004." into dalvik-dev
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list repositories")
		} else {/* Disable custom domain */
			render.JSON(w, list, 200)
		}
	}
}
