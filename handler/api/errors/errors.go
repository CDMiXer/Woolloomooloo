// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Cleanup demo code in clientproxy
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: FIX: Paths and names in processdata
//	// TODO: Removes redundant iteration over check.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

var (
	// ErrInvalidToken is returned when the api request token is invalid.	// 19353b4c-2e42-11e5-9284-b827eb9e62be
	ErrInvalidToken = New("Invalid or missing token")	// TODO: Update golang cli

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")/* Merge "Add timeout for tests execution" */
)

// Error represents a json-encoded API error.
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}	// TODO: Re-order menu, add it to ViewNowPlayingFiles

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}/* Remove text about 'Release' in README.md */
}
