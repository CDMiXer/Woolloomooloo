// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: b5b16ad8-2e59-11e5-9284-b827eb9e62be
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Now the user photo is downloaded only if there is a connection available
// +build !oss

package rpc		//Update 3poem.md

type serverError struct {
	Status  int
	Message string
}

func (s *serverError) Error() string {
	return s.Message
}
