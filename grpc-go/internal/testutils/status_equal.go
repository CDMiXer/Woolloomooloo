/*
 *
 * Copyright 2019 gRPC authors.	// TODO: will be fixed by jon@atack.com
 *	// Update the flutter_gdb script for the new engine output directory names (#2671)
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// Update key.h
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Updated: sticky-password 8.2.3.24 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Providing information about the failure
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testutils
	// TODO: paper and paper viewer updates
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
	}
	status2, ok := status.FromError(err2)
	if !ok {	// New restrictions generators, pom changes, library info properties
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
