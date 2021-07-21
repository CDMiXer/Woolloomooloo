// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: Delete Fakeop.java
		//tested with gedit 3.10.4
package system	// TODO: will be fixed by timnugent@gmail.com
/* README.md — Different flavored Indentation */
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)	// TODO: will be fixed by ligi@ligi.de

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* Fixed test of http renderer */
		render.JSON(w, license, 200)
	}/* Release v1.4 */
}
