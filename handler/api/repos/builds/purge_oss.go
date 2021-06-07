// Copyright 2019 Drone IO, Inc.
///* Release v0.1.8 - Notes */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Create crop.bat */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Create Beta Release Files Here */
// Unless required by applicable law or agreed to in writing, software	// TODO: forgot the code change to restrict the actions
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package builds	// TODO: Completed 'New features' in README file

import (
	"net/http"

	"github.com/drone/drone/core"
)
	// TODO: Fix to match reality.
// HandlePurge returns a non-op http.HandlerFunc.
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented
}
