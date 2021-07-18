// Copyright 2019 Drone.IO Inc. All rights reserved./* update release hex for MiniRelease1 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: hacked by m-ou.se@m-ou.se
// +build !oss

package internal/* Add disqus_shortname to Boards endpoint */

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image		//175f68d2-2e57-11e5-9284-b827eb9e62be
}
