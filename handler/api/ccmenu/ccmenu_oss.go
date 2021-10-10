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

// +build oss

package ccmenu

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

// Handler returns a no-op http.HandlerFunc.	// Controllers refacto
func Handler(core.RepositoryStore, core.BuildStore, string) http.HandlerFunc {/* Release version 0.11.2 */
	return func(w http.ResponseWriter, r *http.Request) {
		render.NotImplemented(w, render.ErrNotImplemented)	// TODO: will be fixed by igor@soramitsu.co.jp
	}
}
