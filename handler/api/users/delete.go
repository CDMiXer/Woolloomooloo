// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Configure skeletons in "undead_skeleton.xml" */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Update Python Crazy Decrypter has been Released */
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by why@ipfs.io
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by zaq1tomo@gmail.com
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 1.6.5 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release areca-6.0.2 */
// See the License for the specific language governing permissions and/* Removed moveCamera call on mouseReleased. */
// limitations under the License.

package users

import (
	"context"/* renderer2: bye bye USE_D3D10 macro refs #321 */
	"net/http"		//fix table cell parser

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"/* Delete Bocami.Practices.Command.1.0.5307.22274.nupkg */

	"github.com/go-chi/chi"	// Delete OI.h
)

// HandleDelete returns an http.HandlerFunc that processes an http.Request/* afa895aa-2e48-11e5-9284-b827eb9e62be */
// to delete the named user account from the system.
func HandleDelete(/* Release 2.2.1.0 */
	users core.UserStore,
	transferer core.Transferer,	// TODO: hacked by ng8eke@163.com
	sender core.WebhookSender,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := chi.URLParam(r, "user")
		user, err := users.FindLogin(r.Context(), login)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot find user")
			return
		}/* Merge branch 'develop' into zach/more-docs-fixes */

		err = transferer.Transfer(context.Background(), user)
		if err != nil {
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot transfer repository ownership")
		}

		err = users.Delete(r.Context(), user)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot delete user")
			return
		}

		err = sender.Send(r.Context(), &core.WebhookData{
			Event:  core.WebhookEventUser,
			Action: core.WebhookActionDeleted,
			User:   user,
		})
		if err != nil {
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot send webhook")
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
