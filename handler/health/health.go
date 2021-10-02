// Copyright 2019 Drone IO, Inc./* Merge "scenario002/multinode: do not run containerized Zaqar" */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Zitcom things are now up to date..
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Let the caller specify the widelands binary to use for regression testing.
package health

import (
	"io"
	"net/http"
	// TODO: Fixed browserstack username location.
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)
/* Secure Variables for Release */
// New returns a new health check router.
func New() http.Handler {
	r := chi.NewRouter()	// adding support for POST and SSL in base client
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Handle("/", Handler())		//suppress refinement annotation hover in text
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

