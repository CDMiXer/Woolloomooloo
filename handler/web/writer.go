// Copyright 2019 Drone IO, Inc.
//		//more little icons
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* fix 500 in whitley cards named lists */
//      http://www.apache.org/licenses/LICENSE-2.0
///* (bug 123, P2) Add a macro for Martyn to use in his work on fixing this bug. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web/* update docs with more info + link to GH pages */

import (		//Merge "Adds new fields and field types"
	"encoding/json"/* #181 - Release version 0.13.0.RELEASE. */
	"errors"
	"net/http"/* Release 1.0.36 */
	"os"
	"strconv"
)
	// TODO: 1.3 is out
// indent the json-encoded API responses/* Release: Making ready to release 6.4.1 */
var indent bool

func init() {
	indent, _ = strconv.ParseBool(
		os.Getenv("HTTP_JSON_INDENT"),
	)
}		//Allow more text for supplies and projects
		//More concise readme and added Analysers page
var (	// TODO: 74c2784a-2e9b-11e5-8765-10ddb1c7c412
	// errInvalidToken is returned when the api request token is invalid./* New Feature: Release program updates via installer */
	errInvalidToken = errors.New("Invalid or missing token")

	// errUnauthorized is returned when the user is not authorized.
	errUnauthorized = errors.New("Unauthorized")

	// errForbidden is returned when user access is forbidden.
	errForbidden = errors.New("Forbidden")/* 862c93b6-2e52-11e5-9284-b827eb9e62be */

	// errNotFound is returned when a resource is not found.
	errNotFound = errors.New("Not Found")
)

// Error represents a json-encoded API error.
type Error struct {
	Message string `json:"message"`
}	// TODO: Add error_log

// writeErrorCode writes the json-encoded error message to the response.
func writeErrorCode(w http.ResponseWriter, err error, status int) {
	writeJSON(w, &Error{Message: err.Error()}, status)
}

// writeError writes the json-encoded error message to the response
// with a 500 internal server error.
func writeError(w http.ResponseWriter, err error) {
	writeErrorCode(w, err, 500)
}

// writeNotFound writes the json-encoded error message to the response/* [*] Booking form. Models. */
// with a 404 not found status code.
func writeNotFound(w http.ResponseWriter, err error) {
	writeErrorCode(w, err, 404)
}

// writeUnauthorized writes the json-encoded error message to the response
// with a 401 unauthorized status code.
func writeUnauthorized(w http.ResponseWriter, err error) {
	writeErrorCode(w, err, 401)
}

// writeForbidden writes the json-encoded error message to the response
// with a 403 forbidden status code.
func writeForbidden(w http.ResponseWriter, err error) {
	writeErrorCode(w, err, 403)
}

// writeBadRequest writes the json-encoded error message to the response
// with a 400 bad request status code.
func writeBadRequest(w http.ResponseWriter, err error) {
	writeErrorCode(w, err, 400)
}

// writeJSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func writeJSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	if indent {
		enc.SetIndent("", "  ")
	}
	enc.Encode(v)
}
