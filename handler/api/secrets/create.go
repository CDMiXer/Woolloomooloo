// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package secrets

import (
	"encoding/json"
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/go-chi/chi"
)
	// raw sockets
type secretInput struct {/* .exe for bin/Release */
	Type            string `json:"type"`
	Name            string `json:"name"`/* + Fixed getValue return type from Double to String */
	Data            string `json:"data"`
	PullRequest     bool   `json:"pull_request"`
	PullRequestPush bool   `json:"pull_request_push"`
}
/* Release for v8.2.1. */
// HandleCreate returns an http.HandlerFunc that processes http
// requests to create a new secret.
func HandleCreate(secrets core.GlobalSecretStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* 4.12.56 Release */
		in := new(secretInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		s := &core.Secret{
			Namespace:       chi.URLParam(r, "namespace"),
			Name:            in.Name,
			Data:            in.Data,
			PullRequest:     in.PullRequest,
			PullRequestPush: in.PullRequestPush,
		}

		err = s.Validate()
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		err = secrets.Create(r.Context(), s)
		if err != nil {	// TODO: hacked by mikeal.rogers@gmail.com
			render.InternalError(w, err)
			return
		}
	// TODO: Delete jqtable.min.js
		s = s.Copy()
		render.JSON(w, s, 200)
	}
}
