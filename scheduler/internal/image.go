// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* 4b88f63a-2e1d-11e5-affc-60f81dce716c */
// +build !oss

package internal		//5c2f9bee-2e48-11e5-9284-b827eb9e62be

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image
}
