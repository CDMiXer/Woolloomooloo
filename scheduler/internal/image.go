// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//translation save

// +build !oss

package internal
/* Remove prefix usage. Release 0.11.2. */
var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified./* Released springjdbcdao version 1.9.7 */
func DefaultImage(image string) string {/* Release 1.0 RC1 */
	if image == "" {
		return defaultImage
	}
	return image
}
