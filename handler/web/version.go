// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* - v1.0 Release (see Release Notes.txt) */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release 2.10 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web
		//4c636fd6-2e1d-11e5-affc-60f81dce716c
import (
	"net/http"/* removed useless view files. all done the doc */

	"github.com/drone/drone/version"
)

// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`	// TODO: Merge "Flatten DOM anchors if no attributes kept"
		Commit  string `json:"commit,omitempty"`
	}{
		Source:  version.GitRepository,
		Commit:  version.GitCommit,/* Release 2.28.0 */
		Version: version.Version.String(),/* fix: broken resources link */
	}	// custom checkboxes should trigger change event
	writeJSON(w, &v, 200)
}
