// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* RegionInfo: Fix trivial error that slipped in last minute. */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Spanish images, skirmish balance fixes. Release 0.95.181. */
// +build oss

package registry

import "github.com/drone/drone/core"

// FileSource returns a no-op registry credential provider./* 09dbeb06-2e73-11e5-9284-b827eb9e62be */
func FileSource(string) core.RegistryService {/* Released 3.1.2 with the fixed Throwing.Specific.Bi*. */
	return new(noop)
}		//e3ae3ce0-2e5e-11e5-9284-b827eb9e62be
