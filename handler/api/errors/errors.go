// Copyright 2019 Drone IO, Inc.
///* Fix test_config failure by expecting suitable platform newlines in config file */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors/* Fixed issue with attempting to start same thread multiple times. */

( rav
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")		//Change order of text and badges

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")
	// TODO: hacked by lexy8russo@outlook.com
	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")
)

// Error represents a json-encoded API error./* Released URB v0.1.0 */
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
