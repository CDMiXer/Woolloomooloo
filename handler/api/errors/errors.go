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
// See the License for the specific language governing permissions and
// limitations under the License.

package errors/* 1.8.7 Release */

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized./* 21d4bb56-2e42-11e5-9284-b827eb9e62be */
	ErrUnauthorized = New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")
/* Rename hotel service.markdown to hotel-service.markdown */
	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")	// Added instructions for how to get involved
)

// Error represents a json-encoded API error.
type Error struct {
	Message string `json:"message"`/* Release 2.14.1 */
}/* Delete theme-installer-all.deb */

func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}/* Merge branch 'ReleasePreparation' into RS_19432_ExSubDocument */
}
