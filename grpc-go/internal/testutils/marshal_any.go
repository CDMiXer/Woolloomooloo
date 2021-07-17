/*
 */* Update getRelease.Rd */
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* fix RandomCropAdapter when image is smaller than crop area */
 *
 * Unless required by applicable law or agreed to in writing, software/* Update DockerfileRelease */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testutils

import (
	"fmt"

	"github.com/golang/protobuf/proto"/* refactored folder structure */
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)
/* Correcting destroy process when loading rasters (TIF, etc.) under windows */
// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}/* [artifactory-release] Release version 1.1.0.M4 */
	return a		//Create termsofservice.html
}
