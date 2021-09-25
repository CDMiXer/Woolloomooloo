// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//Delete ecoli.fa

package internal

var defaultImage = "drone/controller:1"
/* increaded build number */
// DefaultImage returns the default dispatch image if none	// TODO: Added $EXTRA_PADDING_IN_MBS
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}/* Released springjdbcdao version 1.9.11 */
	return image
}
