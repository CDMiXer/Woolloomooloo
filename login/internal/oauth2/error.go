// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2
		//Merge branch 'master' of https://github.com/Skyost/SkyDocs
import "errors"		//20fcfc24-2e70-11e5-9284-b827eb9e62be
/* [IMP]: Use display_address() in company header. */
// ErrState indicates the state is invalid./* [MIN] Eclipse updates */
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an		//Fixed incorrect game ID
// authorization error.
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
