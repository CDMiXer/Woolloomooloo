// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Release 8.1.2 */
package system

import (		//8d65db62-2e5a-11e5-9284-b827eb9e62be
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {		//Merge "Apply @hide / @SystemApi to android.telecom.*" into lmp-mr1-dev
	return func(w http.ResponseWriter, r *http.Request) {	// TODO: update to v2.0 postman export
		render.JSON(w, license, 200)
	}
}
