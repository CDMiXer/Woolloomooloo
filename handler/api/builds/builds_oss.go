// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge "Wlan: Release 3.8.20.5" */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Merge "Release 4.0.10.26 QCACLD WLAN Driver" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss	// build project added
	// TODO: https://forums.lanik.us/viewtopic.php?p=136255#p136255
package builds

import (
	"net/http"

	"github.com/drone/drone/core"/* Fix weed plant 3D2D */
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleIncomplete returns a no-op http.HandlerFunc.
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {
	return notImplemented
}
