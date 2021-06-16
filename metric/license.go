// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Release version 2.9 */
package metric

import "github.com/drone/drone/core"

// License registers the license metrics.
func License(license core.LicenseService) {
	// track days until expires		//CloudApp 3 support documented.
	// track user limit
	// track repo limit
}
