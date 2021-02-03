// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package internal
	// TODO: fixed issue where lat long was not going to server
var defaultImage = "drone/controller:1"
/* Release notes for Trimble.SQLite package */
// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}	// TODO: will be fixed by mikeal.rogers@gmail.com
	return image
}	// TODO: Python path changed
