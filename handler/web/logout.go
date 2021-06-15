// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* bundle-size: 4fa955b792da1432e6e6105f166bb985e29dac72.json */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Update r_conf_cluster.md
package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"/* Release of eeacms/www-devel:21.4.17 */
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.		//FIX method visibility
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")/* Created infrastructure for extended Naming-Service */
		w.Write(	// Create gopro_fix.scad
			dist.MustLookup("/index.html"),/* trigger new build for ruby-head-clang (486f3f4) */
		)
	}/* Merge branch 'master' into ORCIDHUB-31 */
}
