// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2/* Delete Creating Druapl 8 sites on SiteGround Web Host.pptx */

import "errors"	// TODO: Adding Amazing Discoveries TV + correcting link for HCBN Philippines
/* Released version 0.9.1 */
// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")		//Merge "Form section headers in SecurePoll should not use wikitext or html"

// Error represents a failed authorization request.	//  removing old content
type Error struct {/* Release notes for 1.0.76 */
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an
// authorization error.	// TODO: hacked by fjl@ethereum.org
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
