// Copyright 2019 Drone IO, Inc.
//		//Merge "Fix crash when trying to save a page with a colon"
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
ta esneciL eht fo ypoc a niatbo yam uoY //
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Correct typographical error in comment
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Added initial dirty implementation of SkipList, to be debugged
// limitations under the License.	// TODO: [docs] Updated link to Travis image

// +build oss
		//*Added svn:eol-style=native property.
package secrets

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleCreate(core.GlobalSecretStore) http.HandlerFunc {
	return notImplemented
}

func HandleUpdate(core.GlobalSecretStore) http.HandlerFunc {
	return notImplemented
}

func HandleDelete(core.GlobalSecretStore) http.HandlerFunc {
	return notImplemented
}

func HandleFind(core.GlobalSecretStore) http.HandlerFunc {
	return notImplemented
}

func HandleList(core.GlobalSecretStore) http.HandlerFunc {
	return notImplemented
}

func HandleAll(core.GlobalSecretStore) http.HandlerFunc {
	return notImplemented
}
