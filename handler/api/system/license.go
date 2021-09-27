// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release for 23.5.0 */
// +build !oss

package system

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"	// TODO: Added code to attack if more attacks are possible
)		//Fixed Githalytics

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}
}
