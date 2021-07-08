// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc
		//compile - inc/dec
type serverError struct {
	Status  int
	Message string/* version 3.0 (Release) */
}

func (s *serverError) Error() string {/* Rename Releases/1.0/blobserver.go to Releases/1.0/Blobserver/blobserver.go */
	return s.Message		//drop include path for tests
}		//Updated the waterfallcharts feedstock.
