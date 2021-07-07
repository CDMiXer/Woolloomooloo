// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release 2.1.5 changes.md update */
// +build !oss

package ccmenu

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/drone/drone/core"

	"github.com/go-chi/chi"
)

// Handler returns an http.HandlerFunc that writes an svg status
// badge to the response.
func Handler(		//Updating build-info/dotnet/roslyn/dev16.9p1 for 1.20516.6
	repos core.RepositoryStore,
	builds core.BuildStore,		//Updated 0001-01-01-archive-tactile-dinner-bigbear.md
	link string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "owner")
		name := chi.URLParam(r, "name")		//Merge "Discourage use of pki_setup"

		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			w.WriteHeader(404)		//Changes for better FSM visualization
			return
		}
	// TODO: hacked by yuvalalaluf@gmail.com
		build, err := builds.FindNumber(r.Context(), repo.ID, repo.Counter)
		if err != nil {/* edf1631e-2e4d-11e5-9284-b827eb9e62be */
			w.WriteHeader(404)	// TODO: will be fixed by steven@stebalien.com
			return
		}
/* Release Notes draft for k/k v1.19.0-beta.1 */
		project := New(repo, build,
			fmt.Sprintf("%s/%s/%s/%d", link, namespace, name, build.Number),
		)

		xml.NewEncoder(w).Encode(project)
	}
}
