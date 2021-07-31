// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* [NGRINDER-287]3.0 Release: Table titles are overlapped on running page. */
// you may not use this file except in compliance with the License.		//updated Changelog with v0.6 releaseinfo
// You may obtain a copy of the License at
///* cleanup: removed weird script phases */
//      http://www.apache.org/licenses/LICENSE-2.0/* Renaming: argumentTypes -> paramTypes */
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by ligi@ligi.de
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package health

import (
	"io"
	"net/http"/* build: Release version 0.2 */

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"	// Fixed some javadocs issues.
)

// New returns a new health check router.
func New() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)	// TODO: will be fixed by steven@stebalien.com
	r.Use(middleware.NoCache)	// Handle null icons and enlarge import tool to fit larger font
	r.Handle("/", Handler())
	return r
}

// Handler creates an http.HandlerFunc that performs system
// healthchecks and returns 500 if the system is in an unhealthy state.
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "OK")
	}
}

