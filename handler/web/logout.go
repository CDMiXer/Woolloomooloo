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
// See the License for the specific language governing permissions and
// limitations under the License.
		//Incluindo método sleep no objeto rexx
package web

import (/* The entities no longer implement the prototype interface. */
	"net/http"
	// Use Instant instead of LocalDateTime
	"github.com/drone/drone-ui/dist"/* 18c7b4b0-2e5a-11e5-9284-b827eb9e62be */
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")		//Removed "year_id" from Complete Innings query
		w.Write(
			dist.MustLookup("/index.html"),
		)		//add a benchmark folder
	}
}
