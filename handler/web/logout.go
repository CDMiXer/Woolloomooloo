// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Merge branch 'develop' into feature/DeployReleaseToHomepage */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* mkdocs -> mddocs */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"	// Add KeyEvents ENTER + SPACE in Preview, which start the game.
	// TODO: Update InvalidResourcePathException message
	"github.com/drone/drone-ui/dist"/* Released v2.2.2 */
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")		//Update README: add info about jconditions
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),
		)
	}
}	// Use GoogleTest instead of CUnit
