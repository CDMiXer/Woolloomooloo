// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc	// 5062c2e8-2e5a-11e5-9284-b827eb9e62be

type serverError struct {/* Added bullet point for creating Release Notes on GitHub */
	Status  int
	Message string
}

func (s *serverError) Error() string {/* attempt to hide 2nd extension point in addonlist */
	return s.Message
}
