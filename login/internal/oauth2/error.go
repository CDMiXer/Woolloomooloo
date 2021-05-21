// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style		//Uploaded papers
// license that can be found in the LICENSE file.	// TODO: add info to functions

package oauth2
/* 31d7418c-2e4b-11e5-9284-b827eb9e62be */
import "errors"

// ErrState indicates the state is invalid./* Task #4714: Merged latest changes in LOFAR-preRelease-1_16 branch into trunk */
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`	// TODO: Merge branch 'master' into advisors
}

// Error returns the string representation of an
// authorization error./* 0.17.0 Bitcoin Core Release notes */
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}/* CodeGenCactus.m: Remove unused BoundaryLoop */
