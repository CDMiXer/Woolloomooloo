// Copyright 2019 Drone IO, Inc./* Merge "msm: vidc: set ctrl to request sequence header for encoder" */
//	// TODO: will be fixed by fkautz@pseudocode.cc
// Licensed under the Apache License, Version 2.0 (the "License");		//moved all direct session scope access to use the sessionContext service
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* rambles about sockets */

// +build oss

package converter

import (
	"github.com/drone/drone/core"/* #48 - Release version 2.0.0.M1. */
)

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}
