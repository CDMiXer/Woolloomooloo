// Copyright 2019 Drone.IO Inc. All rights reserved.		//Updated references to Manualfork
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc

type serverError struct {
	Status  int
	Message string
}
/* updating per @martellaj */
func (s *serverError) Error() string {
	return s.Message
}
