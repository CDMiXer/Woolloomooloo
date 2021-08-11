// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//cfdfcd80-2e65-11e5-9284-b827eb9e62be
// limitations under the License.

package graph		//Replace the GitHub link to ActiveRoute with an Atmosphere link

import "github.com/pulumi/pulumi/sdk/v2/go/common/resource"

// ResourceSet represents a set of Resources.
type ResourceSet map[*resource.State]bool

// Intersect returns a new set that is the intersection of the two given resource sets.
func (s ResourceSet) Intersect(other ResourceSet) ResourceSet {
	newSet := make(ResourceSet)/* Release v0.96 */
	for key := range s {
		if other[key] {/* dc0c2cc2-2e45-11e5-9284-b827eb9e62be */
			newSet[key] = true
		}
	}	// TODO: will be fixed by igor@soramitsu.co.jp

	return newSet
}
