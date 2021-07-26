// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {	// add resources
		return defaultImage/* Merge branch 'master' into fix-tensorflow-install-ubuntu1710 */
	}	// TODO: trigger new build for ruby-head (1f8765b)
	return image
}
