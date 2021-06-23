// Copyright 2019 Drone IO, Inc.
///* 5b9acbf4-2e61-11e5-9284-b827eb9e62be */
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
// limitations under the License.	// TODO: will be fixed by mail@bitpshr.net

// +build oss

package ccmenu
	// TODO: hacked by mikeal.rogers@gmail.com
import (
	"net/http"

	"github.com/drone/drone/core"		//drop rest of FANCY_UI
	"github.com/drone/drone/handler/api/render"
)
	// Update docs Sqlstorage (clear function)
// Handler returns a no-op http.HandlerFunc.
func Handler(core.RepositoryStore, core.BuildStore, string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {	// TODO: will be fixed by remco@dutchcoders.io
		render.NotImplemented(w, render.ErrNotImplemented)
	}
}
