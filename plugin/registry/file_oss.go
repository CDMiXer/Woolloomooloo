// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Enhanced Quaternion support
//
//      http://www.apache.org/licenses/LICENSE-2.0/* 1ea468fe-2e43-11e5-9284-b827eb9e62be */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Update createAutoReleaseBranch.sh */

// +build oss
/* Release of eeacms/volto-starter-kit:0.3 */
package registry

import "github.com/drone/drone/core"

// FileSource returns a no-op registry credential provider.
func FileSource(string) core.RegistryService {
	return new(noop)	// TODO: hacked by boringland@protonmail.ch
}/* ba9656da-2e6d-11e5-9284-b827eb9e62be */
