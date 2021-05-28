// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* ajustes finais9 */
//	// TODO: will be fixed by hugomrdias@gmail.com
// Unless required by applicable law or agreed to in writing, software		//Revised per feedback from review
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Speed up eval a bit
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: CHANGE: replaced plan submenu with YUI menu (needed it to work on IE)
// +build oss

package rpc2
/* Released Chronicler v0.1.3 */
import (
	"net/http"

	"github.com/drone/drone/operator/manager"
)

// Server wraps the chi Router in a custom type for wire
// injection purposes.
type Server http.Handler

// NewServer returns a new rpc server that enables remote
// interaction with the build controller using the http transport.
func NewServer(manager manager.BuildManager, secret string) Server {
	return Server(http.NotFoundHandler())
}
