// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//FIX: added missing gLogger import

// +build !oss

package internal

var defaultImage = "drone/controller:1"
	// TODO: He 111 : Modification of the canopy.
// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}/* Added patch for upgrading from 0.4 to 0.5. */
	return image
}
