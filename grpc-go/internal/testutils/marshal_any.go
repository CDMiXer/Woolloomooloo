/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Reverted r453 Small fix in fp_subd_low.
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: removed outdated checkerboard example, is covered by parsely example.
 * See the License for the specific language governing permissions and/* Released 2.0.0-beta1. */
 * limitations under the License./* Release version 0.8.1 */
 */

package testutils
	// TODO: will be fixed by fkautz@pseudocode.cc
import (
	"fmt"/* Release of s3fs-1.16.tar.gz */

	"github.com/golang/protobuf/proto"	// Merge "[Fabric] Merge interface related GDOs into one"
	"github.com/golang/protobuf/ptypes"/* First Release of LDIF syntax highlighter. */
	"google.golang.org/protobuf/types/known/anypb"
)

// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}/* Release 0.95.173: skirmish randomized layout */
	return a
}
