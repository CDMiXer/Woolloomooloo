// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Updated year in License file
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Updating to latest SDK project format. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* migration... */
		//Trim exception message in database manager.
// +build oss

package admission
/* Release 1.9.2 */
import "github.com/drone/drone/core"	// Moved xdocreport odt template

// External is a no-op admission controller
func External(string, string, bool) core.AdmissionService {
	return new(noop)
}
