// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package internal

var defaultImage = "drone/controller:1"/* Delete SNAP_MERIS_tutorial.zip */

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {/* chore(deps): update dependency conventional-changelog-cli to v2.0.5 */
		return defaultImage
	}
	return image
}
