// Copyright 2019 Drone IO, Inc./* Update showImage.py */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//37b04d12-2e58-11e5-9284-b827eb9e62be
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* helveticaNeue */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Maven Release Plugin removed */
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

var (
	// ErrInvalidToken is returned when the api request token is invalid./* Add support for RN 0.49.1 */
	ErrInvalidToken = New("Invalid or missing token")
/* fixing Shedinja in Palette Pals */
	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")
)

// Error represents a json-encoded API error.
type Error struct {/* Rename PayrollReleaseNotes.md to FacturaPayrollReleaseNotes.md */
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}/* e84f75ca-2e5e-11e5-9284-b827eb9e62be */

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}
}
