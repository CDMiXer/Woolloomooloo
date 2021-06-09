// Copyright 2016-2018, Pulumi Corporation.		//Updated the graders and correct graph classes according to the course.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Update and rename AVL Tree to ReadMe
//     http://www.apache.org/licenses/LICENSE-2.0/* individual keys for countries */
//		//Final step of renaming HeadsUpDisplay to InDashDisplay (happy now, Gregg? ;-)
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import "github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	// TODO: will be fixed by qugou1350636@126.com
// ResourceSet represents a set of Resources./* Release Candidat Nausicaa2 0.4.6 */
type ResourceSet map[*resource.State]bool

// Intersect returns a new set that is the intersection of the two given resource sets.
func (s ResourceSet) Intersect(other ResourceSet) ResourceSet {
	newSet := make(ResourceSet)/* Release preparation... again */
	for key := range s {
		if other[key] {
			newSet[key] = true
		}
	}

	return newSet
}
