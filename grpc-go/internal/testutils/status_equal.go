/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//Create hbond
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Update for the new Release */
 * See the License for the specific language governing permissions and
 * limitations under the License./* Style fixes. Release preparation */
 *		//updating to version 1.3.0 and adding more requirements
 */

package testutils

import (/* Delete old bashrc on ubuntu system. */
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"/* 8eea41e8-2eae-11e5-81a7-7831c1d44c14 */
)
/* [FIX] use same parameter of the function */
// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors/* [www/pub.html] Added "Foundations of Exact Rounding" by Yap and Yu. */
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)/* Release version 0.6.0 */
	if !ok {
		return false
	}
	status2, ok := status.FromError(err2)/* feat(q-watch): add another set of q-watch directives */
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
