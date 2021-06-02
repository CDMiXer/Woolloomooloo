// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* 7b461b5c-2e4d-11e5-9284-b827eb9e62be */
//      http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/forests-frontend:1.5.4 */
//		//Merge remote-tracking branch 'Wright-Language-Developers/parser' into parser
// Unless required by applicable law or agreed to in writing, software/* Release version 0.2.4 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Cleanup of xtrabackup tests */
package health

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// New returns a new health check router.
func New() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
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

