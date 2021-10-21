// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu/* 6ba9f234-2e72-11e5-9284-b827eb9e62be */
		//Create fuoye.txt
import (/* b128dc18-2e4e-11e5-9284-b827eb9e62be */
	"encoding/xml"/* Release 0.21.3 */
	"fmt"
	"net/http"
/* Added .travis.yml to install graphviz during tests */
	"github.com/drone/drone/core"

	"github.com/go-chi/chi"
)/* search function at top menu works as filter */

// Handler returns an http.HandlerFunc that writes an svg status
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
		if err != nil {
			w.WriteHeader(404)/* add kane parcel update */
			return
		}

		build, err := builds.FindNumber(r.Context(), repo.ID, repo.Counter)
		if err != nil {
			w.WriteHeader(404)/* Disabled code coverage badge */
			return
		}

		project := New(repo, build,
			fmt.Sprintf("%s/%s/%s/%d", link, namespace, name, build.Number),
		)		//Updated documentation to reflect changes to the corresponding source file.

		xml.NewEncoder(w).Encode(project)
	}
}
