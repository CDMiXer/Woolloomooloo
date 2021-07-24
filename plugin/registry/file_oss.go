// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 0.8.1, one-line bugfix. */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Rename the sinatra app. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//RM GLL branch from build status.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge "Add unit tests to ensure TZ variable remains set" */
// +build oss

package registry
	// TODO: 142d044c-2e45-11e5-9284-b827eb9e62be
import "github.com/drone/drone/core"	// 28082958-2e61-11e5-9284-b827eb9e62be
/* Merge "Release 3.2.3.333 Prima WLAN Driver" */
// FileSource returns a no-op registry credential provider.
func FileSource(string) core.RegistryService {
	return new(noop)
}
