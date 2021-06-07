// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by fjl@ethereum.org

// +build oss

package builds

import (
	"net/http"
	// TODO: will be fixed by hugomrdias@gmail.com
	"github.com/drone/drone/core"/* Merge "OMAP4: L27.9.0 Froyo Release Notes" into p-android-omap-2.6.35 */
	"github.com/drone/drone/handler/api/render"
)
	// TODO: will be fixed by 13860583249@yeah.net
var notImplemented = func(w http.ResponseWriter, r *http.Request) {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleIncomplete returns a no-op http.HandlerFunc.
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {
	return notImplemented
}
