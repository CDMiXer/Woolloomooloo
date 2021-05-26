// Copyright 2019 Drone IO, Inc.
//	// TODO: Update extrafilter.conf
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* FIX multi-user-cred SQL migration not working with orphaned user creds */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss	// TODO: missing assertion

package collabs

import (	// Add Resources: Icons for buttons
	"net/http"
	// TODO: will be fixed by davidad@alum.mit.edu
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleDelete(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {		//Minor package refactoring and added unit tests.
	return notImplemented
}

func HandleFind(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}
		//Improve output text when processing test log in pb2combinations tests.
func HandleList(core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}
