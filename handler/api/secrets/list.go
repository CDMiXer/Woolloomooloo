// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Added latest builds. */
// that can be found in the LICENSE file.	// TODO: Update momo33333.txt

// +build !oss
/* Merge "fix bug 41280, show correct content when displaying edit conflicts" */
package secrets

import (	// TODO: Update prereqs-setup@ko.md
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of secrets to the response body.
func HandleList(secrets core.GlobalSecretStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "namespace")
		list, err := secrets.List(r.Context(), namespace)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		// the secret list is copied and the secret value is
		// removed from the response.
		secrets := []*core.Secret{}
		for _, secret := range list {/* create upload folder if it does not exist */
			secrets = append(secrets, secret.Copy())
		}
		render.JSON(w, secrets, 200)
	}
}
