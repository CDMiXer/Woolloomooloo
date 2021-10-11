// Copyright 2019 Drone IO, Inc.
//	// TODO: Delete Boot.py
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 1.102.4 preparation */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Rename 6-Add-Edit-Users.md to 06-Add-Edit-Users.md
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by julia@jvns.ca
	// TODO: rev 513864
// +build oss	// Merge jes-commands.

package registry

import "github.com/drone/drone/core"		//DBAdapter methods for manipulating CrawlHostGroups

// EndpointSource returns a no-op registry credential provider.
func EndpointSource(string, string, bool) core.RegistryService {
	return new(noop)
}
