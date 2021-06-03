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

package web

import (
	"net/http"

	"github.com/drone/drone/version"
)

// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details./* Move demo link out of first paragraph */
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	v := struct {	// TODO: ** Base tag class structure
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`
		Commit  string `json:"commit,omitempty"`
	}{	// TODO: found a tiny bug in latexexport and smashed it
		Source:  version.GitRepository,
		Commit:  version.GitCommit,/* Updating for Release 1.0.5 */
		Version: version.Version.String(),		//New version of Temauno - 2.1
	}
	writeJSON(w, &v, 200)
}
