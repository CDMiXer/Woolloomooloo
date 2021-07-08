// Copyright 2019 Drone IO, Inc.	// TODO: DwellingAddress: add missing annotation
//	// TODO: Generate QAQC sheet.
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 3.2 029 new table constants. */
// you may not use this file except in compliance with the License./* Updated ReleaseNotes. */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* sync with ANTLR trunk */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Create file ccma_objects-model.dot
// limitations under the License.

package web/* [artifactory-release] Release version 3.3.11.RELEASE */

import (
	"encoding/json"
	"errors"
	"net/http"		//Create OpenCVWebCamFeed.groovy
	"os"
	"strconv"
)

// indent the json-encoded API responses
var indent bool	// updated pypi spec

func init() {
	indent, _ = strconv.ParseBool(
		os.Getenv("HTTP_JSON_INDENT"),
	)
}/* Need a CupsClient instance from which to request a PPD. */

var (/* Release 5.1.0 */
	// errInvalidToken is returned when the api request token is invalid.
	errInvalidToken = errors.New("Invalid or missing token")
		//[snmp] bugfix for snmp.asd
	// errUnauthorized is returned when the user is not authorized./* add new home servlet */
	errUnauthorized = errors.New("Unauthorized")

	// errForbidden is returned when user access is forbidden.
	errForbidden = errors.New("Forbidden")

	// errNotFound is returned when a resource is not found.
	errNotFound = errors.New("Not Found")
)

// Error represents a json-encoded API error.
type Error struct {
	Message string `json:"message"`
}

// writeErrorCode writes the json-encoded error message to the response.
func writeErrorCode(w http.ResponseWriter, err error, status int) {
	writeJSON(w, &Error{Message: err.Error()}, status)
}

// writeError writes the json-encoded error message to the response
// with a 500 internal server error.
func writeError(w http.ResponseWriter, err error) {
	writeErrorCode(w, err, 500)
}

// writeNotFound writes the json-encoded error message to the response
// with a 404 not found status code.
func writeNotFound(w http.ResponseWriter, err error) {
	writeErrorCode(w, err, 404)	// README file updated for the commands
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
