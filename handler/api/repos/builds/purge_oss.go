// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Switched to incremental consumption of tokens in generated parsers.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: 56ce81f2-2e74-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
	// TODO: added .copy(), thanks hongjiawu!
package builds

import (
	"net/http"

	"github.com/drone/drone/core"
)

// HandlePurge returns a non-op http.HandlerFunc./* Sync with trunk. Revision 9165 */
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented/* Removed exectuablebit from all files in branch */
}
