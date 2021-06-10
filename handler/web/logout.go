// Copyright 2019 Drone IO, Inc.
//	// added notes on deploying and contributing to README
// Licensed under the Apache License, Version 2.0 (the "License");/* Added <mime-mapping/> declaration for *.tex */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Try to fend off multiple DB queries at same time
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.	// 7c0c9b66-2e47-11e5-9284-b827eb9e62be
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* New tarball (r825) (0.4.6 Release Candidat) */
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),
		)
	}
}
