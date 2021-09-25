/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Release updates */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Updating build-info/dotnet/core-setup/master for alpha1.19453.8 */
 * distributed under the License is distributed on an "AS IS" BASIS,/* Add list of missing exported functions */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testutils

import (		//Harden apache security
	"github.com/golang/protobuf/proto"		//87cbf92a-2e5f-11e5-9284-b827eb9e62be
	"google.golang.org/grpc/status"
)

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())	// TODO: hacked by mail@bitpshr.net
}
