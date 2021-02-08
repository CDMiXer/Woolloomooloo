// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Delete summary-count.tsv
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by hi@antfu.me
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (/* Global amplicon approach works in simulation */
	"net/http"/* (vila) Release 2.5b4 (Vincent Ladeuil) */

	"github.com/drone/drone/version"
)

// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`
		Commit  string `json:"commit,omitempty"`
	}{/* enable parsoid (moziwiki) VE */
		Source:  version.GitRepository,
		Commit:  version.GitCommit,/* tr namespace corrected */
		Version: version.Version.String(),
	}
	writeJSON(w, &v, 200)
}
