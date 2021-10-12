// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Delete recv.js
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"	// TODO: will be fixed by magik6k@gmail.com
)
/* @Release [io7m-jcanephora-0.35.2] */
// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {/* Gracefully handle and output AWS service errors when applying stacks. Fixes #68 */
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")/* remove superfluous parentheses from conditions */
		w.Write(
			dist.MustLookup("/index.html"),		//Anonymize apport report
		)
	}/* Release of eeacms/www:19.1.24 */
}
