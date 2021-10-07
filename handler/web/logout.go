// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//[-bug] fix locale directory substitution in configuration
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// More coding style fixes to autosave manager
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Fixing wrong homepage url
// Unless required by applicable law or agreed to in writing, software/* version bump to v0.4.1 */
// distributed under the License is distributed on an "AS IS" BASIS,		//Update contributing-install-source.rst
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* bc423780-2e52-11e5-9284-b827eb9e62be */

package web/* Version 3.5.2 [KK] */
/* Move button, color, viewport fix */
import (
	"net/http"

	"github.com/drone/drone-ui/dist"
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* fix: force new version test w/ CircleCI + Semantic Release */
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")	// TODO: changed get_absolute_url to use the timezone specified in server settings
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")/* Version 0.957beta released */
		w.Write(
			dist.MustLookup("/index.html"),/* Update eslint to v5.10.0 */
		)
	}
}/* Set root option of rework-npm. Fixes #23 */
