// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: will be fixed by mail@bitpshr.net
// that can be found in the LICENSE file./* Merge "Removing unnecessary call to vp9_setup_interp_filters." */

// +build !oss
	// TODO: Added support for development version of SLEPc. Refactored if-tests
package system

import (
	"net/http"

	"github.com/drone/drone/core"/* feature generator */
	"github.com/drone/drone/handler/api/render"
)

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)/* same for default config */
	}
}
