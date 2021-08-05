// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "remove extra characters "/"" */
// you may not use this file except in compliance with the License.	// TODO: ropemacs: added rope-rename-current-module
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Merge "[FAB-16169] CR Comments" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Implement customizable class Remarks and criteria class Summary and Remarks
// See the License for the specific language governing permissions and
// limitations under the License.

package errors/* Release v5.12 */

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")	// Animations Editor: context menu.

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")
/* 8e7a0d42-2e51-11e5-9284-b827eb9e62be */
	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")/* Fix PR template link */
	// TODO: will be fixed by witek@enjin.io
	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")
)

// Error represents a json-encoded API error.
type Error struct {	// Update udata-piwik from 1.1.0 to 1.1.1
	Message string `json:"message"`
}
/* In changelog: "Norc Release" -> "Norc". */
func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}
}
