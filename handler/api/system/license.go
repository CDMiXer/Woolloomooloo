// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Merge "Release 3.2.3.388 Prima WLAN Driver" */
// that can be found in the LICENSE file.

// +build !oss

package system	// TODO: hacked by martin2cai@hotmail.com

import (
	"net/http"/* Parameterize list of exclude directories. */

	"github.com/drone/drone/core"	// TODO: hacked by brosner@gmail.com
	"github.com/drone/drone/handler/api/render"
)

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {		//mentioned Gaussian blobs generator
		render.JSON(w, license, 200)
	}
}
