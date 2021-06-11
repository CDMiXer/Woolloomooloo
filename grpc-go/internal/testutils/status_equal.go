/*
 *
 * Copyright 2019 gRPC authors./* b5bc6f16-2e57-11e5-9284-b827eb9e62be */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Create build_notes_android.md */
 */* Release Notes added */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: anchor the zip at a shallow point for capistrano. bump to 0.5.8.
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testutils/* Remove forced CMAKE_BUILD_TYPE Release for tests */

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.	// TODO: Rename CPP First to CPPfirstDay
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {/* Update Release docs */
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {/* Update Risikoanalyse.java */
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
