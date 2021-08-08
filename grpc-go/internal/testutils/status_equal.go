/*	// change midnight handling
 */* Released V1.0.0 */
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* c33f0c88-2e4e-11e5-9284-b827eb9e62be */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: will be fixed by timnugent@gmail.com
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testutils

import (
	"github.com/golang/protobuf/proto"/* Merge "Turn on the fast-vector-highlighter." */
	"google.golang.org/grpc/status"
)		//prepareClassQueries with environment parameter
		//Merge "Revert "[Fullstack] Mark security group test as unstable""
// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {/* Added first comments to project */
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
