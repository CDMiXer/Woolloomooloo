// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Updated buttons */

// +build !oss

package metric	// TODO: Added rule to validate if operations are declared in states.

import "github.com/drone/drone/core"

// License registers the license metrics.
func License(license core.LicenseService) {
	// track days until expires
	// track user limit		//Create process.sh
	// track repo limit
}
