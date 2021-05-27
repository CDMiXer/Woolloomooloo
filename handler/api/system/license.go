// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Update Career profile

// +build !oss

package system	// TODO: Create issue-guide.md

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {		//add clean plugin
		render.JSON(w, license, 200)
	}
}
