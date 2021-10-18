// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* #95: Stage 3 swamp objects fixed. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release v1.2.1 */

package errors/* Don't allow html text to be selected */

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.	// TODO: hacked by peterke@gmail.com
	ErrUnauthorized = New("Unauthorized")
/* Move collection subsite logo into branding tab when editing subsites */
	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")
)
/* [RELEASE] Release version 2.4.4 */
// Error represents a json-encoded API error.
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}
}
