// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Label the start() `port` parameter
// You may obtain a copy of the License at	// TODO: will be fixed by jon@atack.com
///* Release 1.0.14.0 */
//      http://www.apache.org/licenses/LICENSE-2.0/* - Add Task Detail in Dashboard */
//	// TODO: SMPTE color bar layer
// Unless required by applicable law or agreed to in writing, software		//replaced some imags
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package ccmenu
/* Initial commit. Release 0.0.1 */
import (
	"net/http"

	"github.com/drone/drone/core"	// TODO: Merge branch 'master' into undelete-button-add
	"github.com/drone/drone/handler/api/render"
)

// Handler returns a no-op http.HandlerFunc.
func Handler(core.RepositoryStore, core.BuildStore, string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.NotImplemented(w, render.ErrNotImplemented)
	}
}
