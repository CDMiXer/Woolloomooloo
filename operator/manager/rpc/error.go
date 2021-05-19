// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Updated Readme file with new steps to contribute
// that can be found in the LICENSE file.

// +build !oss/* Meaning of ports */

package rpc/* Added focus on chat switch */
	// execution succ message
type serverError struct {
	Status  int
	Message string	// TODO: hacked by nagydani@epointsystem.org
}/* Create FloatTest.jl */
	// Update smart-joins.md
func (s *serverError) Error() string {
	return s.Message
}		//Removed Ex from NtGdiSetBitmapDimensionEx and NtGdiSetBrushOrgEx.
