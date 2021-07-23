// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric
/* Release v0.1.0 */
import "github.com/drone/drone/core"
	// Add encryption/decryption in CBC mode
// License registers the license metrics./* rev 560552 */
func License(license core.LicenseService) {
	// track days until expires
	// track user limit
	// track repo limit	// TODO: Solve deprecations
}
