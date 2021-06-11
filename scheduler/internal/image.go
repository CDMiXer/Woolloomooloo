// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* @Release [io7m-jcanephora-0.23.4] */
// that can be found in the LICENSE file.

// +build !oss/* Release version 0.3.0 */

package internal

var defaultImage = "drone/controller:1"/* Add GCD as an example */

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image
}
