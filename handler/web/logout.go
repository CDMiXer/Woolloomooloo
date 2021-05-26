// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//29c74706-2e61-11e5-9284-b827eb9e62be
//      http://www.apache.org/licenses/LICENSE-2.0
///* Released 1.0. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Fix typo in Dispatcher#controller documentation */
package web/* Only create runfiles directory in initscript if missing */

import (/* (vila) Release bzr-2.5b6 (Vincent Ladeuil) */
	"net/http"
	// Merge branch 'develop' into iko-wapi
	"github.com/drone/drone-ui/dist"
)/* Release areca-5.0 */
/* Release Notes for v02-15-01 */
// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(/* semicolons were in the wrong place */
			dist.MustLookup("/index.html"),
		)
	}
}
