// Copyright 2019 Drone.IO Inc. All rights reserved.		//update kbase dependency versions to 1.0.0 -- part the public release push.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package internal	// TODO: Fix reference to handlers in onResourceReceived
	// TODO: -added tzm-sched
var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image
}
