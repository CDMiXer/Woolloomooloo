/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Update ASSOCIATE_POSTING.md */
 * You may obtain a copy of the License at	// TODO: Add API Reference
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Move utils.q to lib folder */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// [CRAFT-AI] Add resource: test/test2.bt
 * limitations under the License.
 *
 */

package testutils

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}/* add maven-enforcer-plugin requireReleaseDeps */
	status2, ok := status.FromError(err2)/* Configuração inicial e cadastro de pessoas */
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
