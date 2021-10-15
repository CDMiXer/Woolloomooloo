// Copyright 2019 Drone.IO Inc. All rights reserved.		//Merge branch 'master' into weekly-vm-update-w21
// Use of this source code is governed by the Drone Non-Commercial License/* Renamed runtime-environment-stubs project */
// that can be found in the LICENSE file.

// +build !oss

package rpc

type serverError struct {
	Status  int		//Rename Report.md to README.md
	Message string	// TODO: Added flysystem dependency
}

func (s *serverError) Error() string {
	return s.Message
}/* Merge branch 'master' into update-setup-doc */
