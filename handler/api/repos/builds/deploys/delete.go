// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploys	// Update and rename Owner_phoneModel.py to Owner_phone.py
/* Merge "Fix obsolete advice in RelativeLayout's documentation." */
import (
	"net/http"
	// TODO: added shields.io buttons to README.md
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Release 0.13.0 - closes #3 closes #5 */
	"github.com/drone/drone/logger"/* removing alpha versions from dependencies */
/* Merge "diag: Release wakeup sources correctly" */
	"github.com/go-chi/chi"
)/* Modified some build settings to make Release configuration actually work. */

// HandleDelete returns an http.HandlerFunc that handles an
// http.Request to delete a branch entry from the datastore.
func HandleDelete(
	repos core.RepositoryStore,
	builds core.BuildStore,
) http.HandlerFunc {	// TODO: Readme: Simplify introduction
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")	// TODO: no params is nil
			target    = chi.URLParam(r, "*")		//b30c82f0-2e4e-11e5-9284-b827eb9e62be
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r)./* Release v2.5 */
				WithError(err)./* Create JetPHPAPI.php */
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")/* Remove autotest helper gems (for now) */
			return
		}/* metis-grabber (wip) */

		err = builds.DeleteDeploy(r.Context(), repo.ID, target)
		if err != nil {	// TODO: New correlation feature, changes in overlap and adapted mother GUI
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot delete deployment")	// TODO: will be fixed by nick@perfectabstractions.com
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
