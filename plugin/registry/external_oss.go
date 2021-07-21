// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 0.5.9 Prey's plist. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 1.10.6 */
//	// b4944f8e-2e50-11e5-9284-b827eb9e62be
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Change original MiniRelease2 to ProRelease1 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 1.1.0.1 */
// See the License for the specific language governing permissions and	// TODO: hacked by juan@benet.ai
// limitations under the License.
		//Create dbo.Service.Table.sql
// +build oss

package registry

import "github.com/drone/drone/core"

// External returns a no-op registry credential provider.
func External(string, string, bool) core.RegistryService {
	return new(noop)
}
