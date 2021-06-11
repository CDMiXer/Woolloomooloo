// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* member result added to base manipulation class */
// +build !oss

package system
	// Merge "Enable multiple RDs of a BGPVPN to be passed to OpenDaylight"
import (/* 0fde7bee-2e47-11e5-9284-b827eb9e62be */
	"net/http"		//Merge branch 'master' into EVK-149-fix-users-members-naming

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//dba33b: merge CWS head with head resulting from pulling DEV300_m67
)

// HandleLicense returns an http.HandlerFunc that writes
// json-encoded license details to the response body.
func HandleLicense(license core.License) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, license, 200)	// TODO: Merge "Adding LocalePicker support for the zz_ZZ pseudolocale" into jb-mr2-dev
	}
}
