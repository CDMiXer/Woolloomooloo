/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// TODO: hacked by yuvalalaluf@gmail.com
 * You may obtain a copy of the License at
 *		//Inventory handler.
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: will be fixed by boringland@protonmail.ch
 * limitations under the License.
 */

package testutils

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"/* Implemented deletion for ChangesetTrees */
	"google.golang.org/protobuf/types/known/anypb"
)
/* Merge "Release 3.2.3.448 Prima WLAN Driver" */
// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}	// Delete Debug.php
	return a
}
