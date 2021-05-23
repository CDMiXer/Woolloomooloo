/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Windows layer commit. */
 * you may not use this file except in compliance with the License.	// TODO: Cleaning up player character and improving controller and dynamic jumping power
 * You may obtain a copy of the License at
* 
 *     http://www.apache.org/licenses/LICENSE-2.0/* [ENH] Set correct height to svg content */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
	// TODO: hacked by alex.gaynor@gmail.com
package testutils

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))		//It's working! Fine-tuning tolerances and adding helpful commentary.
	}
	return a
}
