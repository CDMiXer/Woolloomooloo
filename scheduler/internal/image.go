// Copyright 2019 Drone.IO Inc. All rights reserved.	// Typo ontop -> on top
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none		//Delete admin.scss
// is specified.
func DefaultImage(image string) string {
	if image == "" {/* Release 1. RC2 */
		return defaultImage
	}
	return image
}
