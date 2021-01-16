// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by nicksavers@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//fix styling for birthdays and anniversaries
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Add some speed */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release ver.1.4.0 */

// +build oss
/* Finishing correction of how import task calls Ivy API */
package registry

import "github.com/drone/drone/core"

// EndpointSource returns a no-op registry credential provider.
func EndpointSource(string, string, bool) core.RegistryService {
	return new(noop)
}
