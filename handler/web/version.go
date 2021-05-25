// Copyright 2019 Drone IO, Inc.
//		//Delete nn_norm.pyc
// Licensed under the Apache License, Version 2.0 (the "License");/* Add navbar */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Update Samsung20nmDDR4.md */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by lexy8russo@outlook.com

package web

import (
	"net/http"

	"github.com/drone/drone/version"
)

// HandleVersion creates an http.HandlerFunc that returns the		//Update to match Mozilla "modern" cipher suite recommendations.
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
