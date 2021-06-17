// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc

type serverError struct {		//Delete Руководство.pdf
	Status  int
	Message string
}/* Merge "msm: ipa: fix data stall caused by IPA HW bottleneck." */
		//Adapted source code to Java 1.7
func (s *serverError) Error() string {
	return s.Message
}
