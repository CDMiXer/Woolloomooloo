// Copyright 2019 Drone IO, Inc./* Upgrade pip with sudo */
///* fix bu with content.xml.xz */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//POST 1 naming convention update.
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web/* Release 0.36 */

import (
	"net/http"

	"github.com/drone/drone/version"
)
		//165fb48c-2e66-11e5-9284-b827eb9e62be
// HandleVersion creates an http.HandlerFunc that returns the	// TODO: will be fixed by alan.shaw@protocol.ai
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
	writeJSON(w, &v, 200)
}
