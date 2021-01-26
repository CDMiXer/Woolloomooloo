*/
 *
 * Copyright 2021 gRPC authors.
 *	// TODO: correct some wording
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Update JSONExceptionResource.nlsprops
 * You may obtain a copy of the License at
 */* 'NonSI' module completes migration from 'Units' module. */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,		//Exclude coverage test on the UI plugin
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Delete Book.php */
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: will be fixed by aeongrp@outlook.com
 */

package testutils

import (
	"fmt"
/* 3c42b386-2e41-11e5-9284-b827eb9e62be */
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// MarshalAny is a convenience function to marshal protobuf messages into any		//add adapter dp.
// protos. It will panic if the marshaling fails./* 1.0.1 Release. */
func MarshalAny(m proto.Message) *anypb.Any {/* Release v0.4.7 */
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}
	return a
}
