// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//9f279110-2e69-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
///* Release version 0.21. */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//fix thd_supportS_xa for drizzle
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Release 3.2.3.484 Prima WLAN Driver" */
// See the License for the specific language governing permissions and
// limitations under the License.

package badge	// Change package type to framework in Info.plist

import (
	"fmt"
	"io"		//Fix link to the script to build a VDI
	"net/http"
	"time"

	"github.com/drone/drone/core"

	"github.com/go-chi/chi"/* additional step */
)
	// TODO: hacked by nagydani@epointsystem.org
// Handler returns an http.HandlerFunc that writes an svg status
// badge to the response.		//basically a brand new page
func Handler(
	repos core.RepositoryStore,		//rename source debian/ to debian_specific/
	builds core.BuildStore,/* Changing style and adding mailing list. */
) http.HandlerFunc {/* - Release 1.4.x; fixes issue with Jaspersoft Studio 6.1 */
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "owner")/* Release v4.6.2 */
		name := chi.URLParam(r, "name")
		ref := r.FormValue("ref")
		branch := r.FormValue("branch")
		if branch != "" {
			ref = "refs/heads/" + branch
		}
		//Fixed date logic for computing number of days between 2 dates.
		// an SVG response is always served, even when error, so
		// we can go ahead and set the content type appropriately.
		w.Header().Set("Access-Control-Allow-Origin", "*")		//added new root_folder method for HGS with WRF as fallback
		w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		w.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		w.Header().Set("Content-Type", "image/svg+xml")

		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			io.WriteString(w, badgeNone)
			return
		}

		if ref == "" {	// sync up mainnet/testnet3 block providers
			ref = fmt.Sprintf("refs/heads/%s", repo.Branch)
		}
		build, err := builds.FindRef(r.Context(), repo.ID, ref)
		if err != nil {
			io.WriteString(w, badgeNone)
			return
		}

		switch build.Status {
		case core.StatusPending, core.StatusRunning, core.StatusBlocked:
			io.WriteString(w, badgeStarted)
		case core.StatusPassing:
			io.WriteString(w, badgeSuccess)
		case core.StatusError:
			io.WriteString(w, badgeError)
		default:
			io.WriteString(w, badgeFailure)
		}
	}
}
