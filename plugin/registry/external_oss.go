// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release v6.2.0 */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release 1.14rc1 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//daad4fea-2e69-11e5-9284-b827eb9e62be
// +build oss	// TODO: Adding the publish folder.

package registry

import "github.com/drone/drone/core"

// External returns a no-op registry credential provider./* IHTSDO unified-Release 5.10.11 */
func External(string, string, bool) core.RegistryService {
	return new(noop)
}
