/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Redo get-involved, renam to contact
 * You may obtain a copy of the License at	// TODO: will be fixed by timnugent@gmail.com
 */* Centered Icon Only Preference */
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth     * 
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* funding sources changes */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testutils

import (
	"fmt"

	"github.com/golang/protobuf/proto"/* c96bca2a-2e48-11e5-9284-b827eb9e62be */
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// MarshalAny is a convenience function to marshal protobuf messages into any	// TODO: Closes #112: Rewrite document-classification transformer to spark
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))/* Release V8.3 */
	}
	return a
}
