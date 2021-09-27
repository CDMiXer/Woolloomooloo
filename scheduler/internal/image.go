// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Added some more Pople style basis sets. */

// +build !oss
/* Adding ReleaseNotes.txt to track current release notes. Fixes issue #471. */
package internal
		//Bring README.md up to date
var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none/* add nuget env back */
// is specified.
func DefaultImage(image string) string {/* Typo and grammar fixes in the ActionPack CHANGELOG */
	if image == "" {/* Official Release Version Bump */
		return defaultImage
	}
	return image
}
