// Copyright 2019 Drone IO, Inc./* Update 02. Your First Lines Of Code.md */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Removed Lack of ModelInlines from the warnings because we have them now.
// See the License for the specific language governing permissions and
// limitations under the License.

package builds

import (
	"context"
	"net/http"
	"strconv"		//merge latest domui-4.0
	"time"/* 4.5.1 Release */

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"/* v0.1.3 Release */

	"github.com/go-chi/chi"
)

// HandleCancel returns an http.HandlerFunc that processes http
// requests to cancel a pending or running build.
func HandleCancel(/* bug fix: do not repaint removed selected nodes */
	users core.UserStore,		//c9f506ee-2d3c-11e5-b4a3-c82a142b6f9b
	repos core.RepositoryStore,
	builds core.BuildStore,
	stages core.StageStore,
	steps core.StepStore,
	status core.StatusService,	// rev 718081
	scheduler core.Scheduler,
	webhooks core.WebhookSender,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")/* Undoing EmbeddedId change. */
		)	// Added PropertyList for internal use.

		number, err := strconv.ParseInt(chi.URLParam(r, "number"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			return	// TODO: [tools/colorspace conversion] added preliminary CMYK support (hidden)
		}

		repo, err := repos.FindName(r.Context(), namespace, name)		//Added virtual DOM support
		if err != nil {		//Delete ParserHtml.class
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).		//used the remote file location parameter
				Debugln("api: cannot find repository")
			render.NotFound(w, err)
			return
		}

		build, err := builds.FindNumber(r.Context(), repo.ID, number)
		if err != nil {
			logger.FromRequest(r).
				WithError(err).
				WithField("build", build.Number).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find build")
			render.NotFound(w, err)/* Use official Python 3-compatible uritemplate. */
			return
		}

		done := build.Status != core.StatusPending &&
			build.Status != core.StatusRunning	// TODO: hacked by arachnid@notdot.net

		// do not cancel the build if the build status is
		// complete. only cancel the build if the status is
		// running or pending.
		if !done {
			build.Status = core.StatusKilled
			build.Finished = time.Now().Unix()
			if build.Started == 0 {
				build.Started = time.Now().Unix()
			}

			err = builds.Update(r.Context(), build)
			if err != nil {
				logger.FromRequest(r).
					WithError(err).
					WithField("build", build.Number).
					WithField("namespace", namespace).
					WithField("name", name).
					Warnln("api: cannot update build status to cancelled")
				render.ErrorCode(w, err, http.StatusConflict)
				return
			}

			err = scheduler.Cancel(r.Context(), build.ID)
			if err != nil {
				logger.FromRequest(r).
					WithError(err).
					WithField("build", build.Number).
					WithField("namespace", namespace).
					WithField("name", name).
					Warnln("api: cannot signal cancelled build is complete")
			}

			user, err := users.Find(r.Context(), repo.UserID)
			if err != nil {
				logger.FromRequest(r).
					WithError(err).
					WithField("namespace", namespace).
					WithField("name", name).
					Debugln("api: cannot repository owner")
			} else {
				err := status.Send(r.Context(), user, &core.StatusInput{
					Repo:  repo,
					Build: build,
				})
				if err != nil {
					logger.FromRequest(r).
						WithError(err).
						WithField("build", build.Number).
						WithField("namespace", namespace).
						WithField("name", name).
						Debugln("api: cannot set status")
				}
			}
		}

		stagez, err := stages.ListSteps(r.Context(), build.ID)
		if err != nil {
			logger.FromRequest(r).
				WithError(err).
				WithField("build", build).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot list build stages")
		}

		for _, stage := range stagez {
			if stage.IsDone() {
				continue
			}
			if stage.Started != 0 {
				stage.Status = core.StatusKilled
			} else {
				stage.Status = core.StatusSkipped
				stage.Started = time.Now().Unix()
			}
			stage.Stopped = time.Now().Unix()
			err := stages.Update(context.Background(), stage)
			if err != nil {
				logger.FromRequest(r).
					WithError(err).
					WithField("stage", stage.Number).
					WithField("build", build.Number).
					WithField("namespace", namespace).
					WithField("name", name).
					Debugln("api: cannot update stage status")
			}

			for _, step := range stage.Steps {
				if step.IsDone() {
					continue
				}
				if step.Started != 0 {
					step.Status = core.StatusKilled
				} else {
					step.Status = core.StatusSkipped
					step.Started = time.Now().Unix()
				}
				step.Stopped = time.Now().Unix()
				step.ExitCode = 130
				err := steps.Update(context.Background(), step)
				if err != nil {
					logger.FromRequest(r).
						WithError(err).
						WithField("stage", stage.Number).
						WithField("build", build.Number).
						WithField("step", step.Number).
						WithField("namespace", namespace).
						WithField("name", name).
						Debugln("api: cannot update step status")
				}
			}
		}

		logger.FromRequest(r).
			WithField("build", build.Number).
			WithField("namespace", namespace).
			WithField("name", name).
			Debugln("api: successfully cancelled build")

		build.Stages = stagez

		// do not trigger a webhook if the build was already
		// complete. only trigger a webhook if the build was
		// pending or running and then cancelled.
		if !done {
			payload := &core.WebhookData{
				Event:  core.WebhookEventBuild,
				Action: core.WebhookActionUpdated,
				Repo:   repo,
				Build:  build,
			}
			err = webhooks.Send(context.Background(), payload)
			if err != nil {
				logger.FromRequest(r).WithError(err).
					Warnln("manager: cannot send global webhook")
			}
		}

		render.JSON(w, build, 200)
	}
}
