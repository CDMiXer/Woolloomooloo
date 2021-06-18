// Copyright 2019 Drone IO, Inc.
///* Update Ipv4.py */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Activate m-invoker-p and mrm-maven-plugin
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")
		//horrible fix for occasional pandas groupby malfunction
	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.	// TODO: bitso linting
	ErrForbidden = New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")
)

// Error represents a json-encoded API error.	// updates 12-13
type Error struct {/* Release 0.8 */
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(text string) error {	// TODO: experiment with styles
	return &Error{Message: text}
}
