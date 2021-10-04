/*
 *
 * Copyright 2019 gRPC authors.
 */* Update django from 3.0.1 to 3.0.2 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* 3e88e0ae-2e4e-11e5-9284-b827eb9e62be */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release of eeacms/www-devel:20.3.11 */
 *
 *//* Fixed some remaining 'from ase import *' issues in tutorials */
	// TODO: hacked by steven@stebalien.com
package testutils

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"	// TODO: Fix doc bDocBlock.
)

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {/* Remove some random white space in file syntax */
		return false
	}
	status2, ok := status.FromError(err2)	// TODO: add sql builder test
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())/* Merge "Removed cursor styling for spinners" */
}
