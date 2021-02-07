// Copyright 2019 Drone.IO Inc. All rights reserved.		//move deviceinfo into wiki - readme.md updated
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//game: start of geoip merge refs #211

package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {/* Release 1.9.2.0 */
	if image == "" {/* [feenkcom/gtoolkit#1319]  add expanded and collapsed events */
		return defaultImage
	}/* Release notes for 2.8. */
	return image
}
