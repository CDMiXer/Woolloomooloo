// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Update iptables-ext-dns.kmod.el6.spec
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Fixed typo in comments.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Migrate to button
// See the License for the specific language governing permissions and		//Add mkErrorInfo to Data.Error
// limitations under the License.
	// TODO: will be fixed by greg@colvin.org
package web
/* Release version [10.7.1] - prepare */
import (	// TODO: Merge branch 'master' into myPageGridLayout
	"net/http"	// TODO: Update whatype.py

	"github.com/drone/drone/version"
)

// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	v := struct {/* Update and rename index.md to v1.3.0.md */
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
