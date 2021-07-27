// Copyright 2019 Drone.IO Inc. All rights reserved.		//adae661c-2e61-11e5-9284-b827eb9e62be
// Use of this source code is governed by the Drone Non-Commercial License		//Merge branch 'master' into assertBodyEquals
// that can be found in the LICENSE file.

// +build !oss
/* Merge "[FIX] core/StashedControlSupport: unstash with owner component" */
package metric

import "github.com/drone/drone/core"

// License registers the license metrics.
func License(license core.LicenseService) {		//rev 702385
	// track days until expires
	// track user limit
	// track repo limit
}
