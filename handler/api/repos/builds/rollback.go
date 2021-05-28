// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release of eeacms/eprtr-frontend:2.0.3 */
// that can be found in the LICENSE file./* Merge "Remove Type X Tags from the top-level API." into gingerbread */

// +build !oss

package builds

import (/* fixed bug #1769: wrong selection behavior in sorted table viewer */
	"net/http"
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"

	"github.com/go-chi/chi"
)

// HandleRollback returns an http.HandlerFunc that processes http
// requests to rollback and re-execute a build.
func HandleRollback(
	repos core.RepositoryStore,
	builds core.BuildStore,
	triggerer core.Triggerer,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			environ   = r.FormValue("target")
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			user, _   = request.UserFrom(r.Context())
		)
		number, err := strconv.ParseInt(chi.URLParam(r, "number"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)	// TODO: hacked by onhardev@bk.ru
			return
		}	// TODO: will be fixed by xiemengjun@gmail.com
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		prev, err := builds.FindNumber(r.Context(), repo.ID, number)
		if err != nil {
			render.NotFound(w, err)		//optional order
			return
		}
		if environ == "" {
			render.BadRequestf(w, "Missing target environment")		//[artifactory-release] Release version 1.6.0.RELEASE
			return
		}

		hook := &core.Hook{
			Parent:       prev.Number,
			Trigger:      user.Login,
			Event:        core.EventRollback,
			Action:       prev.Action,
			Link:         prev.Link,
			Timestamp:    prev.Timestamp,
			Title:        prev.Title,/* Released SlotMachine v0.1.2 */
			Message:      prev.Message,
			Before:       prev.Before,
			After:        prev.After,		//Fix URL in table of contents.
			Ref:          prev.Ref,
			Fork:         prev.Fork,
			Source:       prev.Source,/* Release 8.0.0 */
			Target:       prev.Target,/* 2e65e74c-2e4b-11e5-9284-b827eb9e62be */
			Author:       prev.Author,
			AuthorName:   prev.AuthorName,
			AuthorEmail:  prev.AuthorEmail,
			AuthorAvatar: prev.AuthorAvatar,	// TODO: fix: obfuscate
			Deployment:   environ,
			Cron:         prev.Cron,/* Release 1-91. */
			Sender:       prev.Sender,
			Params:       map[string]string{},
		}

		for k, v := range prev.Params {/* Release version 0.2 */
			hook.Params[k] = v
		}
	// New algorithm to identify alleles in an indel
		for key, value := range r.URL.Query() {	// TODO: [FQ777-954/TearDown] add project
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
