// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release v0.0.1-3. */
// +build !oss
	// TODO: fix(DejaMouseDragDropCursor): Add RXJS delay operator
package rpc
/* Merge "Release 3.2.3.488 Prima WLAN Driver" */
type serverError struct {
	Status  int
	Message string	// Merge monthEditor into development
}

func (s *serverError) Error() string {
	return s.Message
}
