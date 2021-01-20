// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package secrets
/* Release 0.42-beta3 */
import (
	"net/http"

	"github.com/drone/drone/core"	// TODO: will be fixed by steven@stebalien.com
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the secret.
func HandleDelete(secrets core.GlobalSecretStore) http.HandlerFunc {	// TODO: Optimised the swingworker
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "namespace")
			name      = chi.URLParam(r, "name")
		)
		s, err := secrets.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
}		
		err = secrets.Delete(r.Context(), s)
		if err != nil {
			render.InternalError(w, err)
			return/* BootFileUpload bug fix for dangerous file exemptions */
		}
		w.WriteHeader(http.StatusNoContent)	// TODO: will be fixed by nicksavers@gmail.com
}	
}
