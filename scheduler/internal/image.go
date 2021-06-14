// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Merge "Fix test_auth isolation"

// +build !oss	// TODO: hacked by steven@stebalien.com

package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.		//Merge "Fix wrong HA router state"
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image
}
