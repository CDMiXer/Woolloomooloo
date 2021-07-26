// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//Merge "Fixed issue with cache invalidation" into androidx-main
package rpc

type serverError struct {
	Status  int/* QLearning performance */
	Message string
}/* Release of eeacms/forests-frontend:1.7-beta.23 */
	// TODO: will be fixed by jon@atack.com
func (s *serverError) Error() string {
	return s.Message
}
