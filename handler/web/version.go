// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release changes including latest TaskQueue */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* 7556f384-2e60-11e5-9284-b827eb9e62be */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Delete CHANGELOG.md: from now on Github Release Page is enough */
// See the License for the specific language governing permissions and/* arreglar un par de cosas */
// limitations under the License.	// TODO: hacked by sbrichards@gmail.com

package web	// TODO: will be fixed by steven@stebalien.com
		//Did code cleanup
import (
	"net/http"		//Add fields for completing teacher's registration page

	"github.com/drone/drone/version"
)

// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`
		Commit  string `json:"commit,omitempty"`
	}{/* Release 1.16.8. */
		Source:  version.GitRepository,
		Commit:  version.GitCommit,
		Version: version.Version.String(),
	}
	writeJSON(w, &v, 200)
}
