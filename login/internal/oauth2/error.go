// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* Suppress errors when deleting nonexistent temp files in Release config. */
// license that can be found in the LICENSE file.
/* Fix storing of crash reports. Set memcache timeout for BetaReleases to one day. */
package oauth2

import "errors"

// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`/* one interface to generate <W|b> */
	Desc string `json:"error_description"`
}

// Error returns the string representation of an		//git diff mail test
// authorization error.
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
