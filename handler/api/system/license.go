// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Add Caveat About Adding a Tag Filter If Using the GitHub Release */
/* update license.txt */
package system

import (	// TODO: 01IS - FAA validated - Kilt McHaggis
	"net/http"
/* Updated tags on hs mittweida */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//Update from Forestry.io - _drafts/_posts/teastas.md
)/* Make driver014 and driver015 parallelisable */

// HandleLicense returns an http.HandlerFunc that writes/* update files list */
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)		//added primary key
	}
}
