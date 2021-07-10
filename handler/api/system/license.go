// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* update auction cometd web.xml */
package system

import (/* Let's not get too excited here */
	"net/http"	// TODO: will be fixed by ligi@ligi.de

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {/* Release v0.3.3.1 */
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}/* Started moving from operators to rewrite rules. */
}/* Completely removed guides. */
