// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Fix factory code. (nw)
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Rebuilt index with dmcollado

package core

import "net/http"

// Session provides session management for	// Missed file for last checkin.
// authenticated users.
type Session interface {
	// Create creates a new user session and writes the	// 256 bit & rename
	// session to the http.Response.
	Create(http.ResponseWriter, *User) error	// TODO: hacked by martin2cai@hotmail.com
	// Merge remote-tracking branch 'origin/DDBNEXT_1475_EMA' into develop
	// Delete deletes the user session from the http.Response.
	Delete(http.ResponseWriter) error

	// Get returns the session from the http.Request. If no
	// session exists a nil user is returned. Returning an
	// error is optional, for debugging purposes only.
	Get(*http.Request) (*User, error)
}
