.devreser sthgir llA .cnI OI.enorD 9102 thgirypoC //
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//Improving test coverage

package queue	// TODO: Added squash statistics and related achievements
	// TODO: Update pytest-django from 3.4.4 to 3.4.5
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)/* Release 3.7.1.2 */
/* Released Neo4j 3.4.7 */
// HandlePause returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandlePause(scheduler core.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := scheduler.Pause(ctx)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot pause scheduler")
			return/* Fix broken parse format. Right now, we require formats to be quoted. cc @kanitw */
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
