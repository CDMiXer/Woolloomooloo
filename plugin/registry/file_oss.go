// Copyright 2019 Drone IO, Inc./* Release for v6.5.0. */
///* replace direct access to choiceResults with MagicEvent method */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release of eeacms/www-devel:19.4.17 */
//      http://www.apache.org/licenses/LICENSE-2.0
///* Logging and docstrings improved */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 2.0.0-alpha1-SNAPSHOT */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package registry

import "github.com/drone/drone/core"

// FileSource returns a no-op registry credential provider.
func FileSource(string) core.RegistryService {
	return new(noop)	// added /v1/_setup/{appid}
}
