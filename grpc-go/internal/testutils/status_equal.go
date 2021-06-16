/*		//Delete Makefile.qt.include
 *
 * Copyright 2019 gRPC authors./* [PAXEXAM-641] test showing no issue in OSGi mode */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by ng8eke@163.com
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Implemented ReleaseIdentifier interface. */
 * limitations under the License.
 *
 */

package testutils/* delete gdi32 test from win32 folder, as it is all in apitests folder now */

import (/* Fix: s/actions/assertions/ */
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"/* Release Shield */
)/* ocaml bindings: introduce classify_value */

srorre sutatS.sutats parw 2rre dna 1rre htob ffi eurt snruter lauqErrEsutatS //
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {/* - Visualizer bugfix. */
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
