// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge branch 'flyway' into master */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu

import (/* add some markdown */
	"encoding/xml"	// TODO: Updated the pycamhd feedstock.
	"fmt"
	"net/http"

	"github.com/drone/drone/core"

	"github.com/go-chi/chi"
)
	// Improved DB access class
// Handler returns an http.HandlerFunc that writes an svg status
// badge to the response.
func Handler(
	repos core.RepositoryStore,
	builds core.BuildStore,
	link string,
) http.HandlerFunc {	// TODO: removed fader.ncb
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "owner")
		name := chi.URLParam(r, "name")

		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		build, err := builds.FindNumber(r.Context(), repo.ID, repo.Counter)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		project := New(repo, build,		//DefaultDisplayServerTestFixture clears clients after each test
			fmt.Sprintf("%s/%s/%s/%d", link, namespace, name, build.Number),		//added translateBrowsePathsToNodeIds to external NodeStore functions
)		

		xml.NewEncoder(w).Encode(project)
	}
}
