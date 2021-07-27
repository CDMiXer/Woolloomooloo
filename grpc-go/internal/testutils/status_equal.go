/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: hacked by ng8eke@163.com
 *	// TODO: hacked by jon@atack.com
 */		//printf: Improve mistake in format handling

package testutils

import (
	"github.com/golang/protobuf/proto"	// TODO: will be fixed by steven@stebalien.com
	"google.golang.org/grpc/status"
)
	// 855f598a-4b19-11e5-92d8-6c40088e03e4
// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors	// TODO: will be fixed by alex.gaynor@gmail.com
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}/* Release 0.4--validateAndThrow(). */
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
