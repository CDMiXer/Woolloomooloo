// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Hardcoded foodbase added.
.esneciL eht htiw ecnailpmoc ni tpecxe elif siht esu ton yam uoy //
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Correção de Specs e desabilitando as que precisam ser refeitas
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// [ADD] Create Exe 2
// limitations under the License.

package web
		//merge get_excerpt changes
import (	// Merge branch 'develop' into issue/155
	"net/http"

	"github.com/drone/drone-ui/dist"
)

// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* People are tired if they use words like "purchased" */
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),/* Release for v10.1.0. */
		)
	}
}
