/*
 */* Release version 1.2.0.M3 */
 * Copyright 2021 gRPC authors.	// Update t1a06-functions-dana.html
 */* php-fpm: set max_children to 6 */
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

package testutils/* Release of eeacms/forests-frontend:2.0-beta.51 */

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"/* Use correct variable. fixes #24014. see #23119. */
)
		//Remove European
// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {/* Release of eeacms/energy-union-frontend:1.7-beta.24 */
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))		//Have no idea
	}
	return a/* Display the E-Tag as well */
}		//refactor: removed global BUILD_FACTORIES
