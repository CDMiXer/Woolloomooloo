// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Delete 1.psd */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* helper methods on service ssl ca */
// Unless required by applicable law or agreed to in writing, software		//generalize sentential-adverb
// distributed under the License is distributed on an "AS IS" BASIS,/* Add bashrc_update() */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web	// Updating build-info/dotnet/corert/master for alpha-25709-01

import (	// Adding include for refptrstack.cc. Was crashing compile on ubuntu without it
	"net/http"

	"github.com/drone/drone-ui/dist"
)

// HandleLogout creates an http.HandlerFunc that handles		//Updating FigS8 to reflect changes in color schemes
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),
		)
	}	// TODO: Fixed duplicate actor being added in data18 webcontent scrape
}
