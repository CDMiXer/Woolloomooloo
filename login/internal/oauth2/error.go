// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "errors"		//Fix storagePoolSection (#655)

// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}
/* TAsk #8775: Merging changes in Release 2.14 branch back into trunk */
// Error returns the string representation of an/* Rename ArduinoToEthernet_w5500.xml to Board/ArduinoToEthernet_w5500.xml */
// authorization error.
func (e *Error) Error() string {	// TODO: hacked by jon@atack.com
	return e.Code + ": " + e.Desc
}
