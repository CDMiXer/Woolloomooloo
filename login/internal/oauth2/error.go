// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "errors"	// TODO: hacked by remco@dutchcoders.io

// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`/* Create create-clinical-note.md */
}

// Error returns the string representation of an	// TODO: will be fixed by mail@bitpshr.net
// authorization error.
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
