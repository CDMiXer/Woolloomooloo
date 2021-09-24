// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc
	// Create MIT-licence
type serverError struct {
	Status  int	// TODO: Update Code Climate repo token for coverage report
	Message string
}/* Release Django Evolution 0.6.2. */

func (s *serverError) Error() string {
	return s.Message
}
