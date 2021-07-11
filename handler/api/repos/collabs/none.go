// Copyright 2019 Drone IO, Inc.
//	// TODO: hacked by igor@soramitsu.co.jp
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Update README.md to include 1.6.4 new Release
// limitations under the License.

// +build oss

package collabs

import (
	"net/http"	// TODO: Merge branch 'master' into component-death
	// stddev and var work as expected.
	"github.com/drone/drone/core"	// Delete dumpdb
	"github.com/drone/drone/handler/api/render"	// TODO: will be fixed by hugomrdias@gmail.com
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)/* Release areca-7.2.13 */
}

func HandleDelete(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}

func HandleFind(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}

func HandleList(core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}
