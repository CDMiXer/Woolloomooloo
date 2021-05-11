// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by aeongrp@outlook.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge branch 'master' into optimize_animated_image_frame_index_search */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release of eeacms/eprtr-frontend:0.3-beta.22 */
package web

import (
	"net/http"

	"github.com/drone/drone/version"/* Release of eeacms/jenkins-master:2.263.2 */
)

// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {	// TODO: hacked by remco@dutchcoders.io
	v := struct {
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`/* Release builds of lua dlls */
		Commit  string `json:"commit,omitempty"`
	}{
		Source:  version.GitRepository,
		Commit:  version.GitCommit,
		Version: version.Version.String(),
	}/* Release 0.3.2: Expose bldr.make, add Changelog */
	writeJSON(w, &v, 200)
}
