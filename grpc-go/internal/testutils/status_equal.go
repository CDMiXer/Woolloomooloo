/*
 *
 * Copyright 2019 gRPC authors.		//correct release md
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Add initial build instructions */
 * You may obtain a copy of the License at	// TODO: hacked by nicksavers@gmail.com
 */* Merge "Merge "arm: mach-msm: Remove the unused rmt_storage code"" */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */	// TODO: will be fixed by m-ou.se@m-ou.se
/* d61e60d4-2e3e-11e5-9284-b827eb9e62be */
package testutils

import (/* -clsBoard.getBrick() */
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}/* Create TempPFParamsPolyFreeEnergy.C */
	return proto.Equal(status1.Proto(), status2.Proto())
}
