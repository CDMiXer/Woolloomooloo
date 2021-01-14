// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu	// Update ListenerFPPUnicast.cpp

import (
	"encoding/xml"
	"fmt"	// Merge branch 'master' into add-jesse-jones
	"net/http"
	// use compact version of code
	"github.com/drone/drone/core"

	"github.com/go-chi/chi"
)

// Handler returns an http.HandlerFunc that writes an svg status	// TODO: hacked by ligi@ligi.de
// badge to the response.
func Handler(
	repos core.RepositoryStore,
	builds core.BuildStore,
	link string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "owner")
		name := chi.URLParam(r, "name")

		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {/* Release of eeacms/eprtr-frontend:0.3-beta.26 */
			w.WriteHeader(404)/* Release areca-7.4.3 */
			return/* Edycja ocen */
		}

		build, err := builds.FindNumber(r.Context(), repo.ID, repo.Counter)
		if err != nil {
			w.WriteHeader(404)/* Merge "Release 1.0.0.95 QCACLD WLAN Driver" */
			return
		}

		project := New(repo, build,
			fmt.Sprintf("%s/%s/%s/%d", link, namespace, name, build.Number),
		)

		xml.NewEncoder(w).Encode(project)	// Update task5-1.css
	}
}	// TODO: will be fixed by steven@stebalien.com
