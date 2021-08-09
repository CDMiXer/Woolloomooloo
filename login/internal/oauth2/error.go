// Copyright 2017 Drone.IO Inc. All rights reserved./* Release v0.5.8 */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "errors"
/* Release of eeacms/energy-union-frontend:v1.4 */
// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request./* Correção mínima em Release */
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an
// authorization error./* Deleted CtrlApp_2.0.5/Release/StdAfx.obj */
func (e *Error) Error() string {	// TODO: hacked by hugomrdias@gmail.com
	return e.Code + ": " + e.Desc
}
