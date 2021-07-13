// Copyright 2019 Drone IO, Inc./* Merge branch 'master' into greenkeeper/less-3.0.0 */
///* fix for when tabbar controller present */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Create installed.txt
// distributed under the License is distributed on an "AS IS" BASIS,	// Add AboutActivity for simple building "About" dialog.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Fix Python 3. Release 0.9.2 */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package registry

import "github.com/drone/drone/core"

// EndpointSource returns a no-op registry credential provider.
func EndpointSource(string, string, bool) core.RegistryService {
	return new(noop)
}
