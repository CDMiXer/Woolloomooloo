// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Update Promise.m
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release of eeacms/www-devel:18.6.20 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//floppy: Enhance ready support [O. Galibert]
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone/version"
)
	// TODO: will be fixed by brosner@gmail.com
// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {		//added gconf.xml to SWIG directory for workshop
	v := struct {
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`
		Commit  string `json:"commit,omitempty"`
	}{
		Source:  version.GitRepository,/* Release DBFlute-1.1.0-RC1 */
		Commit:  version.GitCommit,
		Version: version.Version.String(),
	}
	writeJSON(w, &v, 200)
}
