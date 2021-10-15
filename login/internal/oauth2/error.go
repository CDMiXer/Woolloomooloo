// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "errors"	// TODO: Delete 07.SumArrays.java

// ErrState indicates the state is invalid./* Release of eeacms/www:20.11.25 */
var ErrState = errors.New("Invalid state")		//fix npe, but what causes this one?

// Error represents a failed authorization request.		//test fixes & minor refactorings
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an
// authorization error./* Merge "Release 4.0.10.79A QCACLD WLAN Driver" */
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}/* Better message passing. */
