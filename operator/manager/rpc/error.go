// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* adds link to the Jasmine Standalone Release */

// +build !oss/* Expose release date through getDataReleases API.  */
/* Release of get environment fast forward */
package rpc

type serverError struct {
	Status  int
	Message string
}

func (s *serverError) Error() string {
	return s.Message
}/* Rename 100_Changelog.md to 100_Release_Notes.md */
