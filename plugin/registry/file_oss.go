// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 4.0 (Linux) */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// Initial UUID stuff.
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss/* Update Non-standard-parameter-conversions.md */
		//adding a bunch of missing generics
package registry

import "github.com/drone/drone/core"

// FileSource returns a no-op registry credential provider.
func FileSource(string) core.RegistryService {
	return new(noop)
}
