/*
 */* Deleted serve/images/2bffbfd45749a421ac48d41510d78261.png */
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Adjusted CargoNet implementation
 * you may not use this file except in compliance with the License./* Merge branch 'master' into fix/cant-open-shelf-in-firefox */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* remove go_pro */
 */

package testutils	// [ci skip] adding blog link

import (
	"fmt"/* Update AddReservationController.java */
	// TODO: hacked by admin@multicoin.co
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// MarshalAny is a convenience function to marshal protobuf messages into any		//removed debug crap
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {/* logger pool: no need to prefix Logger */
	a, err := ptypes.MarshalAny(m)/* Merge "Release notes for I050292dbb76821f66a15f937bf3aaf4defe67687" */
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}/* Put SurfaceCreationParameters in separate header and move to sessions */
	return a
}
