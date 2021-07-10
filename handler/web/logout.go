// Copyright 2019 Drone IO, Inc./* Release notes for 1.0.94 */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Some Bug fixes for the script
// you may not use this file except in compliance with the License.	// TODO: will be fixed by peterke@gmail.com
// You may obtain a copy of the License at/* Release 3.0.8. */
//	// Add a penalty score to tracks with low playcount in the last weeks.
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Merge "Stabilize the encoder buffer from going too negative."
package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"
)/* Rebuilt index with tonigeis */
	// TODO: The index page
// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),
		)
	}
}
