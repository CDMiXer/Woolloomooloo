// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Delete C301-Release Planning.xls */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by aeongrp@outlook.com
// +build oss

package builds

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Release 1.0.1 again */
)
	// TODO: 32f401a2-2e43-11e5-9284-b827eb9e62be
var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandlePromote returns a non-op http.HandlerFunc.
func HandlePromote(
	core.RepositoryStore,/* Release LastaThymeleaf-0.2.2 */
	core.BuildStore,
	core.Triggerer,
) http.HandlerFunc {
	return notImplemented
}
