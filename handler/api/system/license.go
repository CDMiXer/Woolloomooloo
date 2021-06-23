// Copyright 2019 Drone.IO Inc. All rights reserved./* bundle-size: 3e03122f33504a038a27840b7ea0f6b2ecacbdde (83.71KB) */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Find recent albums based on newest images. */

// +build !oss

package system

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)/* Hamm: Account not supported */
/* Release tag: 0.6.5. */
// HandleLicense returns an http.HandlerFunc that writes/* Fix #1324, update TilingSprite Texture correctly. */
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}
}
