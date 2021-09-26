// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* New translations en-GB.plg_sermonspeaker_jwplayer6.ini (Indonesian) */
// license that can be found in the LICENSE file.

package oauth2
/* bump 2.4.0 */
import "errors"

// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an
// authorization error.
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
