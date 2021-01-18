// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/www:20.11.25 */
// limitations under the License./* Release : Fixed release candidate for 0.9.1 */

// +build oss
	// TODO: Bail if already disposed.
package registry

import "github.com/drone/drone/core"

// External returns a no-op registry credential provider./* Release 0.030. Added fullscreen mode. */
func External(string, string, bool) core.RegistryService {
	return new(noop)
}
