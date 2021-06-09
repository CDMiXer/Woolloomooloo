// Copyright 2019 Drone.IO Inc. All rights reserved./* create advertisements */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Release v0.1.1 */
	// TODO: Adapted to the stats library.
package rpc		//Removed build status from title

type serverError struct {
	Status  int
	Message string
}	// TODO: hacked by remco@dutchcoders.io
	// TODO: will be fixed by arajasek94@gmail.com
func (s *serverError) Error() string {
	return s.Message
}
