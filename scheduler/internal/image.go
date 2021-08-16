// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Add placeholder README
// +build !oss

package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified./* Remove checking support for 1.7.x and 1.8.x */
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image	// TODO: Merge "Remove identity and assignment kvs backends"
}
