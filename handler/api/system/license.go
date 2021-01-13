// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: will be fixed by alan.shaw@protocol.ai

package system

import (
	"net/http"

	"github.com/drone/drone/core"/* Merge "Release 3.2.3.338 Prima WLAN Driver" */
	"github.com/drone/drone/handler/api/render"
)/* Ignore errors from Sentry when the application is offline */

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.		//Remove remains from the old bookmarklet code
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}
}
