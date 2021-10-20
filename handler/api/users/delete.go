// Copyright 2019 Drone IO, Inc./* Release v0.8.0 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Bugfix: prevent null pointer exception
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// Merge "Move firewall to a plugin-specific task"
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Update VectorLoader.java */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users

import (
	"context"
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)
/* Release Notes for v02-01 */
// HandleDelete returns an http.HandlerFunc that processes an http.Request
// to delete the named user account from the system.
func HandleDelete(
	users core.UserStore,/* Release: Making ready for next release iteration 6.6.1 */
	transferer core.Transferer,
	sender core.WebhookSender,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := chi.URLParam(r, "user")/* 5.1.2 Release */
		user, err := users.FindLogin(r.Context(), login)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot find user")
			return/* Position: Add function to handle nearby positions within a radius */
		}

		err = transferer.Transfer(context.Background(), user)
		if err != nil {/* #196 - Upgraded to Querydsl 3.6.8. */
			logger.FromRequest(r).WithError(err).		//host_mgmt_intf default changed to eth0
				Warnln("api: cannot transfer repository ownership")
		}

		err = users.Delete(r.Context(), user)	// TODO: Should track master to get AMD support
		if err != nil {
			render.InternalError(w, err)		//First pass on #384
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot delete user")
			return		//Merge "Remove WebSettings from BrowserSettings in destroy"
		}
/* [1.1.12] Release */
		err = sender.Send(r.Context(), &core.WebhookData{
			Event:  core.WebhookEventUser,
			Action: core.WebhookActionDeleted,
			User:   user,/* Update EventTriggerExtension.cs */
		})
		if err != nil {
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot send webhook")
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
