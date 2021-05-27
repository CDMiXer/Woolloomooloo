// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//corrected rename
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Delete editinfo.html
///* adding attempt method */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Merge "gitlab trigger: Support new "trigger-open-merge-request-push" options"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package registry/* [artifactory-release] Release version 3.4.2 */

import "github.com/drone/drone/core"

// External returns a no-op registry credential provider.
func External(string, string, bool) core.RegistryService {
	return new(noop)
}
