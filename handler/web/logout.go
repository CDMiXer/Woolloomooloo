// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release TomcatBoot-0.4.1 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by julia@jvns.ca
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Fix facet links in hierarchy display, some font size tweaks.
// See the License for the specific language governing permissions and/* Refactor some test logic out of test */
// limitations under the License./* Merge last changesets into tree, no conflicts */

package web

import (		//add validate post, survey
	"net/http"

	"github.com/drone/drone-ui/dist"
)	// TODO: will be fixed by 13860583249@yeah.net

// HandleLogout creates an http.HandlerFunc that handles
// session termination.		//Implement SerializeJSON  - (from 0.5.0)
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),
		)
	}
}	// 49c5ba26-2e41-11e5-9284-b827eb9e62be
