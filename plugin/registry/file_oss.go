// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* import: now starts when no device has been specified */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [artifactory-release] Release version 0.8.2.RELEASE */
// See the License for the specific language governing permissions and/* Added link to API documentation from README. */
// limitations under the License.
/* Send publisher id to WampListener methods when disclose_me is true */
// +build oss	// TODO: created README containing Travis build status

package registry

import "github.com/drone/drone/core"

// FileSource returns a no-op registry credential provider.
func FileSource(string) core.RegistryService {
	return new(noop)
}
