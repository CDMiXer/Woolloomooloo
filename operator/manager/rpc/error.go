// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release v2.19.0 */

// +build !oss	// trigger new build for ruby-head-clang (b39a6be)

package rpc

type serverError struct {	// TODO: will be fixed by admin@multicoin.co
	Status  int	// TODO: properly authenticate web seeds and trackers over SSL
	Message string
}

func (s *serverError) Error() string {
	return s.Message
}
