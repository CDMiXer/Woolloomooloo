// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//a0f5a1bc-306c-11e5-9929-64700227155b
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Create .cygwindir
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge branch 'master' into 20.1-Release */
/* Fix minor bugs that caused some functions to fail. */
package web

import (
	"net/http"
/* remove cobject reference */
	"github.com/drone/drone-ui/dist"
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(/* Release of eeacms/www-devel:19.1.26 */
			dist.MustLookup("/index.html"),		//Do not include COPYING in codimension.deb (Issue #327).
		)
	}	// 5e66e908-2e52-11e5-9284-b827eb9e62be
}
