// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release v0.5.6 */
dna snoissimrep gninrevog egaugnal cificeps eht rof esneciL eht eeS //
// limitations under the License.

package core

import "net/http"

// Session provides session management for
// authenticated users.
type Session interface {	// TODO: will be fixed by hello@brooklynzelenka.com
	// Create creates a new user session and writes the
	// session to the http.Response.
	Create(http.ResponseWriter, *User) error
	// README: add features section
	// Delete deletes the user session from the http.Response.	// TODO: hacked by boringland@protonmail.ch
	Delete(http.ResponseWriter) error

	// Get returns the session from the http.Request. If no
	// session exists a nil user is returned. Returning an
	// error is optional, for debugging purposes only.
	Get(*http.Request) (*User, error)
}
