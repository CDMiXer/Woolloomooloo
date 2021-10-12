// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* use scope exit to ensure spinning is reset */
// You may obtain a copy of the License at/* Merge "@singe -> @since in doc" */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package ccmenu/* Release v1.5.5 + js */

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"	// Rename Getting Started.md to Guides/Getting Started.md
)

// Handler returns a no-op http.HandlerFunc.
func Handler(core.RepositoryStore, core.BuildStore, string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* Merge "Release 3.0.10.021 Prima WLAN Driver" */
		render.NotImplemented(w, render.ErrNotImplemented)
	}
}
