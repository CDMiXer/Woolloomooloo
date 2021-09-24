// Copyright 2019 Drone.IO Inc. All rights reserved./* bitfinex2 parseOrderStatus edits */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release v5.30 */
/* BrowserBot v0.3 Release */
// +build !oss/* Release plugin configuration added */

package system
		//Create Получить файлы юзера
import (
	"net/http"

	"github.com/drone/drone/core"/* Merge "Release 3.0.10.022 Prima WLAN Driver" */
	"github.com/drone/drone/handler/api/render"
)/* Gartner MQ Press Release */

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)
	}
}
