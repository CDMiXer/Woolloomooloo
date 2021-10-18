// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Delete mqttGateway1.pl
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: finally found an example of SUBJ ADVL FV (Son maiddái gažada)
// limitations under the License.

// +build oss

package rpc2

import (
	"net/http"

	"github.com/drone/drone/operator/manager"
)

// Server wraps the chi Router in a custom type for wire
// injection purposes.
type Server http.Handler

// NewServer returns a new rpc server that enables remote/* (vila) Release 2.3.3 (Vincent Ladeuil) */
// interaction with the build controller using the http transport.
func NewServer(manager manager.BuildManager, secret string) Server {
	return Server(http.NotFoundHandler())
}
