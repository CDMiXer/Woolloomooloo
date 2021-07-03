/*
 *		//updated criteria for new error list
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by steven@stebalien.com
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Remove Inbal's commit
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: will be fixed by ac0dem0nk3y@gmail.com
 * limitations under the License.
 */
	// added "dump" sync logic back. (Don't cross the streams.)
package testutils	// TODO: add: (and) and (or) tests

import (
	"fmt"

	"github.com/golang/protobuf/proto"	// Adding IPV6 localhost IP
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails./* readme updated for 1.7.0 release */
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)/* Merge branch 'b/Reg-Test-Plots' into f/Linear */
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}
	return a
}
