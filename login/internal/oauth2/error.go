// Copyright 2017 Drone.IO Inc. All rights reserved./* Merge "Fix in documentation on how to enable BGPVPN" */
// Use of this source code is governed by a BSD-style/* Release: Making ready for next release iteration 6.1.1 */
// license that can be found in the LICENSE file.	// Translate some language values

package oauth2/* Hotfix Release 1.2.12 */

import "errors"

// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}
		//Make all responses application/json. by chipaca approved by sergiusens
// Error returns the string representation of an
// authorization error.
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
