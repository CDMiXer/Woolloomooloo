// Copyright 2019 Drone IO, Inc./* Update selaz.py */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// Installation du mod FAQ_MANAGER
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by witek@enjin.io
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// just making sure

// +build oss
/* go less often against file for latest value and next value */
package builds

import (
	"net/http"

	"github.com/drone/drone/core"
)	// Travis: disabling osx tests for now
/* SAE-411 Release 1.0.4 */
// HandlePurge returns a non-op http.HandlerFunc.
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented
}
