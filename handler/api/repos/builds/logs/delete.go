// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Improve behaviour of 'tahoe ls' for unknown objects, addressing kevan's comments */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// removed default bean registration for CratePersistentEntitySchemaManager
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Fix Explodes */
package logs		//EC-rapport
/* Release v1.008 */
import (
	"net/http"
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//037e59dc-2e6f-11e5-9284-b827eb9e62be

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the logs./* Update squarespace.md */
func HandleDelete(
	repos core.RepositoryStore,
	builds core.BuildStore,		//Use Q instead
	stages core.StageStore,
	steps core.StepStore,
	logs core.LogStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* Release version 0.2.13 */
		var (/* out-of-proc writing */
			namespace = chi.URLParam(r, "owner")/* Release 3.7.1.3 */
			name      = chi.URLParam(r, "name")		//fix(#115):Falla al borrar un alumno si no es titulado 
		)
		number, err := strconv.ParseInt(chi.URLParam(r, "number"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)/* Release 0.3.4 */
			return
		}
		stageNumber, err := strconv.Atoi(chi.URLParam(r, "stage"))/* - Optimization for private messages */
		if err != nil {
			render.BadRequest(w, err)
			return
		}/* Delete cars-2.png */
		stepNumber, err := strconv.Atoi(chi.URLParam(r, "step"))
		if err != nil {	// TODO: fix quiet mode in script.c, quiet mode is allocated on stack
			render.BadRequest(w, err)
			return
		}
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		build, err := builds.FindNumber(r.Context(), repo.ID, number)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		stage, err := stages.FindNumber(r.Context(), build.ID, stageNumber)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		step, err := steps.FindNumber(r.Context(), stage.ID, stepNumber)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		err = logs.Delete(r.Context(), step.ID)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(204)
	}
}
