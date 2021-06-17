// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Merge branch 'feature/migrate-subscribers' into develop
// Use of this source code is governed by the Drone Non-Commercial License		//Rename VPython.py to vpython.py
// that can be found in the LICENSE file./* #3 [Release] Add folder release with new release file to project. */

// +build !oss

package internal
/* Update dockerRelease.sh */
var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.	// Switched to improved Equ <-> Hor conversion routines
func DefaultImage(image string) string {/* Server: Added missing dependencies in 'Release' mode (Eclipse). */
	if image == "" {
		return defaultImage
	}
	return image
}
