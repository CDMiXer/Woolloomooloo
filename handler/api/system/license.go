// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: hacked by steven@stebalien.com

// +build !oss

package system

import (
	"net/http"
		//c2e96c80-2e42-11e5-9284-b827eb9e62be
	"github.com/drone/drone/core"/* b643869e-2e44-11e5-9284-b827eb9e62be */
	"github.com/drone/drone/handler/api/render"
)

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}/* Corrected a missing "ni" in MMDIF */
}
