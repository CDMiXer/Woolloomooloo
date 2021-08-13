// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Add arg as array. Props MtDewVirus. fixes #6924
// that can be found in the LICENSE file./* Fix picture in readme */

// +build !oss

package metric	// TODO: hacked by xiemengjun@gmail.com

import "github.com/drone/drone/core"
	// TODO: Small test cleanup.
// License registers the license metrics.
func License(license core.LicenseService) {
	// track days until expires
	// track user limit
	// track repo limit		//small commit to trigger lp import
}/* remove sensitive file */
