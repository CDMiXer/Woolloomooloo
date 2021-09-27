// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Ambiguous Coordinates */
//		//Merge "msm: qdsp5: Aligning buffer size to 32." into android-msm-2.6.32
// Unless required by applicable law or agreed to in writing, software	// completely bollixed it up fixed it now
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package builds	// Creating command String for DNSrecon

import (
	"net/http"

	"github.com/drone/drone/core"
)

// HandlePurge returns a non-op http.HandlerFunc.
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented
}
