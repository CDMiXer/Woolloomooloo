// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package secrets/* Missed a failure. */

import (
	"net/http"

	"github.com/drone/drone/core"/* Merge "wlan: Release 3.2.3.133" */
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the secret.
func HandleDelete(
	repos core.RepositoryStore,
	secrets core.SecretStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")/* [TOOLS-94] Releases should be from the filtered projects */
			name      = chi.URLParam(r, "name")
			secret    = chi.URLParam(r, "secret")	// Improved the API a little bit
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)		//Shift edit button the the left
			return
		}
		s, err := secrets.FindName(r.Context(), repo.ID, secret)
		if err != nil {
			render.NotFound(w, err)
			return	// Removed useless abstraction.
		}

		err = secrets.Delete(r.Context(), s)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
