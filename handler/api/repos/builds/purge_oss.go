// Copyright 2019 Drone IO, Inc.
///* replace deprecated env.parse() with .parsePublic() and env.parseInternal() */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Cleaning Up. Getting Ready for 1.1 Release */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Allow plugin reload from console */
// limitations under the License.

// +build oss

package builds

import (
	"net/http"

	"github.com/drone/drone/core"		//Merge "Changes configuration_ref to configuration"
)

// HandlePurge returns a non-op http.HandlerFunc.		//Updated Prenat to new ISML syntax
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented
}
