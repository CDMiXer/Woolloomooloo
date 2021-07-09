// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "errors"
		//Slight changes for new SASS layouts
.dilavni si etats eht setacidni etatSrrE //
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request./* v1.1.14 Release */
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an
// authorization error.
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
