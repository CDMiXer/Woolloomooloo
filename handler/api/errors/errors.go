// Copyright 2019 Drone IO, Inc.
///* [artifactory-release] Release version 3.1.0.RC2 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Eggdrop v1.8.3 Release Candidate 1 */
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")		//dplay: support for premium content

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")
/* Release branch */
	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")
)

// Error represents a json-encoded API error./* Released springjdbcdao version 1.8.12 */
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message./* Transport addressing refactoring */
func New(text string) error {
	return &Error{Message: text}
}
