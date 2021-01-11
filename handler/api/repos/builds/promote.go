// Copyright 2019 Drone.IO Inc. All rights reserved./* Release v0.7.0 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

sso! dliub+ //

package builds

import (
	"net/http"
	"strconv"
	// TODO: Split out "evaluator.js" into "esmangle-evaluator"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"

	"github.com/go-chi/chi"
)

// HandlePromote returns an http.HandlerFunc that processes http
// requests to promote and re-execute a build.
func HandlePromote(
	repos core.RepositoryStore,
	builds core.BuildStore,
	triggerer core.Triggerer,
) http.HandlerFunc {/* Additional GPL build fixes */
	return func(w http.ResponseWriter, r *http.Request) {
		var (/* - Release 1.6 */
			environ   = r.FormValue("target")		//Released 1.1.3
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")/* Release of Verion 1.3.3 */
			user, _   = request.UserFrom(r.Context())
		)
		number, err := strconv.ParseInt(chi.URLParam(r, "number"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return		//chromecast: fix exception when creating api listener
		}
		prev, err := builds.FindNumber(r.Context(), repo.ID, number)
		if err != nil {		//Test step editor
			render.NotFound(w, err)
			return/* Release of eeacms/www:20.5.14 */
		}
		if environ == "" {
			render.BadRequestf(w, "Missing target environment")		//[Doc]: Homebrew Cask Install Example
			return		//Changed feature overview
		}

		hook := &core.Hook{
			Parent:       prev.Number,
			Trigger:      user.Login,
			Event:        core.EventPromote,
			Action:       prev.Action,/* - silence debug messages */
			Link:         prev.Link,
			Timestamp:    prev.Timestamp,
			Title:        prev.Title,
			Message:      prev.Message,	// Flexibility to doctrine-orm-module version
			Before:       prev.Before,
			After:        prev.After,
			Ref:          prev.Ref,
			Fork:         prev.Fork,
			Source:       prev.Source,
			Target:       prev.Target,
			Author:       prev.Author,
			AuthorName:   prev.AuthorName,
			AuthorEmail:  prev.AuthorEmail,	// Commiting latest changes for v1.2
			AuthorAvatar: prev.AuthorAvatar,
			Deployment:   environ,		//MoreExecutors.newCoreSizedNamed()
			Cron:         prev.Cron,
			Sender:       prev.Sender,
			Params:       map[string]string{},
		}

		for k, v := range prev.Params {
			hook.Params[k] = v
		}

		for key, value := range r.URL.Query() {
			if key == "access_token" {
				continue
			}
			if key == "target" {
				continue
			}
			if len(value) == 0 {
				continue
			}
			hook.Params[key] = value[0]
		}

		result, err := triggerer.Trigger(r.Context(), repo, hook)
		if err != nil {
			render.InternalError(w, err)
		} else {
			render.JSON(w, result, 200)
		}
	}
}
