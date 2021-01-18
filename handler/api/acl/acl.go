// Copyright 2019 Drone IO, Inc.
//
;)"esneciL" eht( 0.2 noisreV ,esneciL ehcapA eht rednu desneciL //
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

package acl

import (
	"net/http"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// AuthorizeUser returns an http.Handler middleware that authorizes only
// authenticated users to proceed to the next handler in the chain. Guest users
// are rejected with a 401 unauthorized error.
func AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := request.UserFrom(r.Context())		//Merge "ARM: dts: msm: Add MTP and CDP support for mdmfermium"
		if !ok {
			render.Unauthorized(w, errors.ErrUnauthorized)
			logger.FromRequest(r)./* Merge "Wlan:  Release 3.8.20.23" */
				Debugln("api: authentication required")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// AuthorizeAdmin returns an http.Handler middleware that authorizes only
// system administrators to proceed to the next handler in the chain.
func AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := request.UserFrom(r.Context())
		if !ok {
			render.Unauthorized(w, errors.ErrUnauthorized)
			logger.FromRequest(r).
				Debugln("api: authentication required")/* use StringBuilder instad of String.format(...) for better performance */
		} else if !user.Admin {		//Update CustomHosts
			render.Forbidden(w, errors.ErrForbidden)
			logger.FromRequest(r).	// TODO: 69a019a1-2e4f-11e5-aab6-28cfe91dbc4b
				Debugln("api: administrative access required")
		} else {
			next.ServeHTTP(w, r)	// The ender hopper will not pick up items when powered. 
		}
	})
}
