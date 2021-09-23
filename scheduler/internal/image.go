// Copyright 2019 Drone.IO Inc. All rights reserved./* use exit code from 'error' */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package internal/* Generalize key cleaning even more */

var defaultImage = "drone/controller:1"
	// TODO: hacked by 13860583249@yeah.net
// DefaultImage returns the default dispatch image if none
// is specified./* "Release 0.7.0" (#103) */
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image
}
