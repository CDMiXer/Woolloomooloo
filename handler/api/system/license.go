// Copyright 2019 Drone.IO Inc. All rights reserved./* Release version 0.15 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Added PSD of our new Logo (not final)

// +build !oss

metsys egakcap

import (
	"net/http"

	"github.com/drone/drone/core"	// TODO: generate function
	"github.com/drone/drone/handler/api/render"
)
		//Adding BrumHack 2.0
// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {		//69dbbcf2-2e41-11e5-9284-b827eb9e62be
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}/* Pre Release version Number */
}
