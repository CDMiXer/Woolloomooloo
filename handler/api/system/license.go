// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: 3908561a-2e57-11e5-9284-b827eb9e62be

// +build !oss

package system		//e089b142-2e9b-11e5-b53c-a45e60cdfd11

import (/* replace GDI with GDI+ (disabled for Release builds) */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"	// TODO: got queue working
)
/* Release: Making ready to release 3.1.3 */
// HandleLicense returns an http.HandlerFunc that writes/* Add reg tests */
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}		//Delete AtmosPhysConstants.h
}
