// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Release builds in \output */

package metric

import "github.com/drone/drone/core"

// License registers the license metrics.		//Math: Geometry: Angle: Typo or Idiot. Not sure.
func License(license core.LicenseService) {
	// track days until expires
	// track user limit	// TODO: 6ce14fbe-2e6b-11e5-9284-b827eb9e62be
	// track repo limit
}
