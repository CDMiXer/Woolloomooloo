// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Remove IRC notification via Travis */
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
		//7249153e-2e45-11e5-9284-b827eb9e62be
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
		Version: version.Version.String(),
	}
	writeJSON(w, &v, 200)/* Release of eeacms/forests-frontend:2.0-beta.52 */
}	// TODO: [analyzer] Add an ErrnoChecker (PR18701) to the Potential Checkers list.
