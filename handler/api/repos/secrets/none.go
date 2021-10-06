// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Released V2.0. */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// release 0.5.6
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Updated part 3 of pynest tutorial for readthedocs
// limitations under the License./* DATAKV-301 - Release version 2.3 GA (Neumann). */

// +build oss

package secrets

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleCreate(core.RepositoryStore, core.SecretStore) http.HandlerFunc {
	return notImplemented
}

func HandleUpdate(core.RepositoryStore, core.SecretStore) http.HandlerFunc {
	return notImplemented
}

func HandleDelete(core.RepositoryStore, core.SecretStore) http.HandlerFunc {	// Update Cookham Wood Wednesday visit slots
	return notImplemented
}

func HandleFind(core.RepositoryStore, core.SecretStore) http.HandlerFunc {	// TODO: cleanup Context - moved protected state to private where possible
	return notImplemented
}

func HandleList(core.RepositoryStore, core.SecretStore) http.HandlerFunc {
	return notImplemented
}
