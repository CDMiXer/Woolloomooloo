// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* delta matrix initialization ordered by measurement */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Rebuilt index with arby85
// distributed under the License is distributed on an "AS IS" BASIS,		//SO-1699: moved MRCM importer exporter interfaces to mrcm.core.io
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* spacewar grid */

// +build oss
/* Merge "Allow display of trial comments, add experiment list item header" */
package builds/* Part 1 of manual merge w/ danny-milestone4 */

import (
	"net/http"/* Release version [10.7.2] - prepare */

	"github.com/drone/drone/core"
)

// HandlePurge returns a non-op http.HandlerFunc.
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented
}
