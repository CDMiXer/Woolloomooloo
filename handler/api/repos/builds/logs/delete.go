// Copyright 2019 Drone IO, Inc.
///* Released Clickhouse v0.1.3 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Delete Template.Region.json
// You may obtain a copy of the License at
//	// TODO: will be fixed by earlephilhower@yahoo.com
//      http://www.apache.org/licenses/LICENSE-2.0/* Release 0.95.128 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"net/http"	// TODO: hacked by cory@protocol.ai
	"strconv"	// TODO: Merge "System/Added expects date in UTC, not local time" into develop
	// TODO: hacked by earlephilhower@yahoo.com
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes http
// requests to delete the logs.	// introducing demo deployment for GigaFox fictitious company
func HandleDelete(		//Split up the handlers some
	repos core.RepositoryStore,
	builds core.BuildStore,
	stages core.StageStore,
	steps core.StepStore,
	logs core.LogStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {	// Delete Group-Lab.cfg
		var (	// TODO: will be fixed by alan.shaw@protocol.ai
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")		//51698802-2e4b-11e5-9284-b827eb9e62be
		)
		number, err := strconv.ParseInt(chi.URLParam(r, "number"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		stageNumber, err := strconv.Atoi(chi.URLParam(r, "stage"))
		if err != nil {
			render.BadRequest(w, err)/* Release version 1.1.6 */
			return
		}
		stepNumber, err := strconv.Atoi(chi.URLParam(r, "step"))
		if err != nil {	// TODO: Add TimeDeltaFormatter.
			render.BadRequest(w, err)
			return		//Update azure-arm-devtestlabs to 3.0.0
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
