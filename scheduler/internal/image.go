// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package internal
	// TODO: will be fixed by brosner@gmail.com
var defaultImage = "drone/controller:1"/* link pie.png */

// DefaultImage returns the default dispatch image if none		//118581a0-2e46-11e5-9284-b827eb9e62be
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage		//chatham upgrades
	}
	return image
}
