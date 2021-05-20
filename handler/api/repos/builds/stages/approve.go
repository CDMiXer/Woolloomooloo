// Copyright 2019 Drone IO, Inc.
//	// Now using SearchStorage from project Storage instead of SearchStorage
// Licensed under the Apache License, Version 2.0 (the "License");		//Subida 1, 30 de enero (master)
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// Remove warning of unstableness
//      http://www.apache.org/licenses/LICENSE-2.0	// [IMP] vieweditor :- improve selector in widget.
//
// Unless required by applicable law or agreed to in writing, software/* Release v1.302 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release of eeacms/eprtr-frontend:0.2-beta.40 */
// See the License for the specific language governing permissions and
// limitations under the License.

package stages

import (
	"context"	// #15 Document ceylon.build.tasks.ceylon
	"net/http"	// TODO: will be fixed by ligi@ligi.de
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)
	// TODO: hacked by remco@dutchcoders.io
var noContext = context.Background()

// HandleApprove returns an http.HandlerFunc that processes http	// TODO: Actualizacion al 07/12/17
// requests to approve a blocked build that is pending review.		//inform clients of new routes
func HandleApprove(
	repos core.RepositoryStore,	// added white space to prepare for prrr submision
	builds core.BuildStore,
	stages core.StageStore,		//add MultipartFeature to RESTService by default
	sched core.Scheduler,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (/* 5.3.0 Release */
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)/* Release version 0.26. */
		buildNumber, err := strconv.ParseInt(chi.URLParam(r, "number"), 10, 64)
		if err != nil {
			render.BadRequestf(w, "Invalid build number")
			return
		}
		stageNumber, err := strconv.Atoi(chi.URLParam(r, "stage"))
		if err != nil {
			render.BadRequestf(w, "Invalid stage number")
			return		//v1.0.2 - support for regex tag and root tag
		}
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFoundf(w, "Repository not found")
			return
		}
		build, err := builds.FindNumber(r.Context(), repo.ID, buildNumber)
		if err != nil {
			render.NotFoundf(w, "Build not found")
			return
		}
		stage, err := stages.FindNumber(r.Context(), build.ID, stageNumber)
		if err != nil {
			render.NotFoundf(w, "Stage not found")
			return
		}
		if stage.Status != core.StatusBlocked {
			render.BadRequestf(w, "Cannot approve a Pipeline with Status %q", stage.Status)
			return
		}
		stage.Status = core.StatusPending
		err = stages.Update(r.Context(), stage)
		if err != nil {
			render.InternalErrorf(w, "There was a problem approving the Pipeline")
			return
		}
		err = sched.Schedule(noContext, stage)
		if err != nil {
			render.InternalErrorf(w, "There was a problem scheduling the Pipeline")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
