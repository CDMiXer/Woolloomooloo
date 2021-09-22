// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Delete mph_zpool_hashrefinery_bench_start.bat
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by joshua@yottadb.com

package web		//Merge branch 'develop' into event-fixes

import (		//creation working
	"net/http"

	"github.com/drone/drone-ui/dist"
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {	// TODO: Punch IE styles.
	return func(w http.ResponseWriter, r *http.Request) {	// TODO: will be fixed by davidad@alum.mit.edu
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),/* set height of tourlist header  */
		)
	}
}
