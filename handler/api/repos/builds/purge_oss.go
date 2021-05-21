// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Add Dark Sky Attribution
// you may not use this file except in compliance with the License.	// TODO: will be fixed by admin@multicoin.co
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//improve box text position
// +build oss

package builds

import (
	"net/http"	// Removed name from package details.

	"github.com/drone/drone/core"
)

// HandlePurge returns a non-op http.HandlerFunc.
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented	// TODO: hacked by steven@stebalien.com
}
