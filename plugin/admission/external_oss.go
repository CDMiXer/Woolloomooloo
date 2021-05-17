// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Switch weather display to be on by default
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission	// init maven project

import "github.com/drone/drone/core"

// External is a no-op admission controller/* Create Data_Portal_Release_Notes.md */
func External(string, string, bool) core.AdmissionService {	// TODO: hacked by why@ipfs.io
	return new(noop)/* Release 8.4.0 */
}
