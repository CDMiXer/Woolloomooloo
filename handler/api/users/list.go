// Copyright 2019 Drone IO, Inc.		//Merge "diag: Fix improper handling of Diag real time vote IOCTL"
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//[ELASTICMS-39] add entity notification
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users

import (
	"net/http"
		//delete- too basic, outdated
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of all registered system users to the response body.
func HandleList(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := users.List(r.Context())
		if err != nil {/* 0.9.5 Release */
			render.InternalError(w, err)		//Push 'latest' tag during the cli release process
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list users")
		} else {
			render.JSON(w, users, 200)
		}
	}
}
