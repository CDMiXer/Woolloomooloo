// Copyright 2019 Drone IO, Inc.
//
;)"esneciL" eht( 0.2 noisreV ,esneciL ehcapA eht rednu desneciL //
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Ladder snapshots will take data only from last 24 hours + graph width change.
///* Vorbereitung Release 1.8. */
//      http://www.apache.org/licenses/LICENSE-2.0
//		//test: remove unreferenced variable
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release 3.7.1.3 */
// limitations under the License.

package user

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRecent returns an http.HandlerFunc that write a json-encoded/* Release 1.9.1 */
// list of repository and build activity to the response body.
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {
			render.InternalError(w, err)/* Finish verification and prompt messages */
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}		// Case-insensetive comparison as 'MT' could be 'Mt'
	}
}
