// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//rename world panel/parameters to dimension panel/parameters

package system

import (
	"net/http"/* Release notes: Delete read models */

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Release of eeacms/eprtr-frontend:0.4-beta.1 */
)/* Corrigido pequeno problema de compilador */

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* Fix bug reported by scrutinizer */
		render.JSON(w, license, 200)		//Update and rename 0817.md to 0823.md
	}
}		//Für neue TaxStatement Version aufgebohrt
