// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Merged Lastest Release */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

var (		//Significant readme update
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")	// TODO: hacked by witek@enjin.io

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")
)

// Error represents a json-encoded API error.
type Error struct {
	Message string `json:"message"`	// TODO: Update example.py to use flask.ext compatibility imports.
}

func (e *Error) Error() string {
	return e.Message/* HomeTimeLineFragment: Save cache when pause */
}

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}/* Beta Release Version */
}		//Remove kytos dependency from dev.in
