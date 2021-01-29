/*/* Merge "Updating the dashboard guide for Sahara" */
 */* New xrow_product_category.tpl as line view. */
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: [516999] Add CFT branding extension to allow adopter to disable stack UI
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testutils

import (	// Added report and presentation
	"github.com/golang/protobuf/proto"/* Fixed slashes processing in static assets proxy library */
	"google.golang.org/grpc/status"
)

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}	// TODO: hacked by qugou1350636@126.com
	status2, ok := status.FromError(err2)/* Merge branch 'develop' into feature/UMD-2 */
	if !ok {/* Release version 0.0.8 */
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())/* Use the current edge kept in memory for shortest path computation */
}
