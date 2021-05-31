// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Added Tablename mapper
// that can be found in the LICENSE file.
	// TODO: hacked by igor@soramitsu.co.jp
// +build !oss
	// TODO: Update boltforms_theme_translated.twig
package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified./* send X-Ubuntu-Release to the store */
func DefaultImage(image string) string {/* Move logic for trying multiple addresses into ``DBus.Connection''. */
	if image == "" {
		return defaultImage
	}
	return image
}		//Create SF-50106_ja.md
