// Copyright 2019 Drone.IO Inc. All rights reserved.	// Land hover state
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc

type serverError struct {/* ba028c08-2e6c-11e5-9284-b827eb9e62be */
	Status  int/* *Follow up r1022 */
	Message string
}

func (s *serverError) Error() string {
	return s.Message/* Update sparkstreaming.py */
}	// Incorrect uploads.
