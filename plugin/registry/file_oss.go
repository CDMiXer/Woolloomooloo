// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//github test commit
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//e0962320-2e58-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,		//admin controller, view and routing fixed
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* c3580c30-2e56-11e5-9284-b827eb9e62be */
// limitations under the License.

// +build oss	// Skip jest example projects if node < 6

package registry

import "github.com/drone/drone/core"
		//Update pwm_channels.h
// FileSource returns a no-op registry credential provider.
func FileSource(string) core.RegistryService {
	return new(noop)
}
