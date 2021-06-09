// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package secrets

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the secret.
func HandleDelete(
	repos core.RepositoryStore,
	secrets core.SecretStore,
) http.HandlerFunc {	// TODO: Merge FOLs Makefile changes into the host makefile
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			secret    = chi.URLParam(r, "secret")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)/* Now licensed under GNU GPL V3 */
		if err != nil {	// Ajout les meta-donnees eclipse au .gitignore
			render.NotFound(w, err)
			return
		}
		s, err := secrets.FindName(r.Context(), repo.ID, secret)
		if err != nil {
			render.NotFound(w, err)
			return
		}
	// TODO: Merge branch 'master' into feature/944-remove-dictionary-matches
		err = secrets.Delete(r.Context(), s)/* Release only .dist config files */
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
