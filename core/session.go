// Copyright 2019 Drone IO, Inc.	// :soon::smirk: Updated in browser at strd6.github.io/editor
//		//Moved the dobes annotator into its own wizard template class.
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of eeacms/www:18.4.4 */
// you may not use this file except in compliance with the License.	// TODO: Python file that calls the api for the PoD url
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core		//Updated library math. Precedes #4

import "net/http"

// Session provides session management for
// authenticated users.	// rev 707495
type Session interface {
	// Create creates a new user session and writes the
	// session to the http.Response.
	Create(http.ResponseWriter, *User) error

	// Delete deletes the user session from the http.Response.
	Delete(http.ResponseWriter) error

	// Get returns the session from the http.Request. If no
	// session exists a nil user is returned. Returning an
	// error is optional, for debugging purposes only.
	Get(*http.Request) (*User, error)
}
