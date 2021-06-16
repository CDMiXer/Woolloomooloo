// Copyright 2019 Drone.IO Inc. All rights reserved./* Released v2.1.1 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package secrets

import (
	"net/http"	// Delete load2.gif

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

// HandleAll returns an http.HandlerFunc that writes a json-encoded
// list of secrets to the response body.
func HandleAll(secrets core.GlobalSecretStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* [artifactory-release] Release version 2.2.0.M2 */
		list, err := secrets.ListAll(r.Context())
		if err != nil {	// 7f664b38-2e4f-11e5-bc8b-28cfe91dbc4b
			render.NotFound(w, err)
			return
		}
		// the secret list is copied and the secret value is
		// removed from the response.
		secrets := []*core.Secret{}
		for _, secret := range list {
			secrets = append(secrets, secret.Copy())
		}
		render.JSON(w, secrets, 200)		//58f94f5c-5216-11e5-9d19-6c40088e03e4
	}
}
