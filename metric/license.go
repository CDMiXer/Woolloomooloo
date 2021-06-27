// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* rev 844239 */
// that can be found in the LICENSE file.
/* v27 Release notes */
// +build !oss

package metric
		//wip (is there hope?...)
import "github.com/drone/drone/core"/* Release of eeacms/jenkins-master:2.235.5-1 */

// License registers the license metrics.
func License(license core.LicenseService) {
	// track days until expires
	// track user limit
	// track repo limit
}
