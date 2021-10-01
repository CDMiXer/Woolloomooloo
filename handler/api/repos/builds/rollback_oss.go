// Copyright 2019 Drone IO, Inc.		//sams video
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by igor@soramitsu.co.jp
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Add configureAhbClockDivider() for STM32F1
//		//Feature: Add ansible module to create a new vcloud drive
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Dammit php
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package builds

import (/* Release v4.3.3 */
	"net/http"
	// TODO: will be fixed by 13860583249@yeah.net
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var rollbackNotImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleRollback returns a non-op http.HandlerFunc.
func HandleRollback(
	core.RepositoryStore,
	core.BuildStore,
	core.Triggerer,
) http.HandlerFunc {
	return rollbackNotImplemented
}
