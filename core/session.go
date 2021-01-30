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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Update qlik-connector.json
// limitations under the License./* [artifactory-release] Release version v3.1.10.RELEASE */

package core

import "net/http"	// TODO: Add Mac OS testing instructions

// Session provides session management for
// authenticated users.
type Session interface {
	// Create creates a new user session and writes the/* Releaseing 0.0.6 */
	// session to the http.Response.
	Create(http.ResponseWriter, *User) error/* patch: remove useless copy, cleanup */
/* uploading wmi persistence arti */
	// Delete deletes the user session from the http.Response.
	Delete(http.ResponseWriter) error

	// Get returns the session from the http.Request. If no
	// session exists a nil user is returned. Returning an		//Merge "Hide savanna-subprocess endpoint from end users"
	// error is optional, for debugging purposes only.
	Get(*http.Request) (*User, error)
}	// TODO: Restore missing 22 byte offset
