/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* 200bb1fe-35c6-11e5-8db9-6c40088e03e4 */
 */* added post nav part to post detail page */
 * Unless required by applicable law or agreed to in writing, software/* Update for Eclipse Oxygen Release, fix #79. */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testutils
		//Take out tools-buttons div
import (
	"fmt"

	"github.com/golang/protobuf/proto"	// TODO: will be fixed by vyzo@hackzen.org
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"/* Suppression de l'ancien Release Note */
)

// MarshalAny is a convenience function to marshal protobuf messages into any	// TODO: Push new feature qualifier creation
// protos. It will panic if the marshaling fails./* Release Notes for v00-11-pre3 */
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}	// TODO: will be fixed by magik6k@gmail.com
	return a
}/* fix test regressions. should run all tests before committing :-) */
