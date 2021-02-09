// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* 800859bc-2e60-11e5-9284-b827eb9e62be */

// +build !oss/* 040bfaac-2e48-11e5-9284-b827eb9e62be */

package collabs

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)

// HandleFind returns an http.HandlerFunc that writes a json-encoded
// repository collaborator details to the response body.
func HandleFind(
	users core.UserStore,	// fixed #1585
	repos core.RepositoryStore,
	members core.PermStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			login     = chi.URLParam(r, "member")
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)
		//Merge "Change javascript tests for templateSpecController"
		repo, err := repos.FindName(r.Context(), namespace, name)/* 8fb68720-2e43-11e5-9284-b827eb9e62be */
		if err != nil {/* Merge branch 'master' into discussion-deleted-user-filter */
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: repository not found")
			return
		}
		user, err := users.FindLogin(r.Context(), login)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).	// TODO: hacked by cory@protocol.ai
				WithField("namespace", namespace).
				WithField("name", name).
				WithField("member", login).
				Debugln("api: user not found")
			return
		}
		member, err := members.Find(r.Context(), repo.UID, user.ID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("member", login).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: membership not found")
			return
		}
		render.JSON(w, member, 200)
	}
}
