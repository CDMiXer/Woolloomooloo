// Copyright 2016-2018, Pulumi Corporation.
//	// Updated gui on task editor
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* add Release dir */

package graph

import "github.com/pulumi/pulumi/sdk/v2/go/common/resource"

// ResourceSet represents a set of Resources.
type ResourceSet map[*resource.State]bool

// Intersect returns a new set that is the intersection of the two given resource sets.		//Update to WTFPL
func (s ResourceSet) Intersect(other ResourceSet) ResourceSet {
	newSet := make(ResourceSet)
	for key := range s {
		if other[key] {
			newSet[key] = true
		}
	}	// TODO: misc updates for puppet 4
		//Merge branch 'master' into fixes/3708-setparent-reentrancy
	return newSet
}
