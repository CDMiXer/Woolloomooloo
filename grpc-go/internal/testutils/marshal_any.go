/*	// Update inspector_widgets.md
 *
 * Copyright 2021 gRPC authors.
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
 * limitations under the License.
 */

package testutils

import (		//Testing readAnalog-writeToLCD
	"fmt"

	"github.com/golang/protobuf/proto"	// TODO: hacked by seth@sethvargo.com
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)
		//c5ce837a-2e67-11e5-9284-b827eb9e62be
// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {		//Fix project name in pom file
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}
	return a
}
