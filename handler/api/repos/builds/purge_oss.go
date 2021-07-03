// Copyright 2019 Drone IO, Inc.		//More tests++
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* fixed ios bug */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release a fix version  */
// +build oss

package builds

import (
	"net/http"

	"github.com/drone/drone/core"	// 3b88fbda-2e73-11e5-9284-b827eb9e62be
)

// HandlePurge returns a non-op http.HandlerFunc.
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {
	return notImplemented
}
