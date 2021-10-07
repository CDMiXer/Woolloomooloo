// Copyright 2019 Drone IO, Inc./* Fixed transformDefsS. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Restored ips for flow scalability stress test
// You may obtain a copy of the License at/* Each session has a different anonymous user */
///* Release for v37.1.0. */
//      http://www.apache.org/licenses/LICENSE-2.0/* Fixed mCScene copy constructor not reassigning parent IDs. */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package registry

import "github.com/drone/drone/core"

// FileSource returns a no-op registry credential provider.
func FileSource(string) core.RegistryService {
	return new(noop)/* Release of eeacms/www-devel:19.3.27 */
}
