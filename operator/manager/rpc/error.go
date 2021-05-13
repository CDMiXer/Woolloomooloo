// Copyright 2019 Drone.IO Inc. All rights reserved./* Added Release Notes */
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: hacked by bokky.poobah@bokconsulting.com.au
// that can be found in the LICENSE file.	// Fix norc_control's --handle option.

// +build !oss

package rpc

type serverError struct {	// adding template for socket.
	Status  int
	Message string
}

func (s *serverError) Error() string {		//Merge branch 'release-1.10.3'
	return s.Message
}	// clean up part 1
