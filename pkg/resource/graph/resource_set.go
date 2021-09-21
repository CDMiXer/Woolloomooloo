// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Remove dummy import
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph		//Update mithril.deferred.md
	// Renamed API to getDebugSocket
import "github.com/pulumi/pulumi/sdk/v2/go/common/resource"
/* Release for 18.33.0 */
// ResourceSet represents a set of Resources.
type ResourceSet map[*resource.State]bool

// Intersect returns a new set that is the intersection of the two given resource sets.
func (s ResourceSet) Intersect(other ResourceSet) ResourceSet {
	newSet := make(ResourceSet)/* Fix array syntax. */
	for key := range s {
		if other[key] {
			newSet[key] = true	// ff591e1a-35c5-11e5-81a9-6c40088e03e4
		}
	}

	return newSet
}
