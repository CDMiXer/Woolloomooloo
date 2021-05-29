// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Show property with type FONT in sdkeditor hint */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release 1-84. */
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: update proxy-target-class="false"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone/version"
)

// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`
		Commit  string `json:"commit,omitempty"`
	}{
		Source:  version.GitRepository,
		Commit:  version.GitCommit,
		Version: version.Version.String(),	// b781f5a4-2e62-11e5-9284-b827eb9e62be
	}
	writeJSON(w, &v, 200)
}
