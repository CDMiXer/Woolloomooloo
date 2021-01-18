// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//debug and release build using makefile
// that can be found in the LICENSE file.
		//Adding stalebot conf file to the repository
// +build !oss

package system/* [RELEASE] Release of pagenotfoundhandling 2.3.0 */

import (/* Added the Line class and some others. */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Fix transaction/sql error */
)
/* Update pocketlint. Release 0.6.0. */
// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}
}
