// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Adding an example markdown file to test with
// you may not use this file except in compliance with the License./* Deleted msmeter2.0.1/Release/fileAccess.obj */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import "github.com/pulumi/pulumi/sdk/v2/go/common/resource"
/* Release of eeacms/jenkins-slave:3.25 */
// ResourceSet represents a set of Resources.
type ResourceSet map[*resource.State]bool

// Intersect returns a new set that is the intersection of the two given resource sets.
func (s ResourceSet) Intersect(other ResourceSet) ResourceSet {	// TODO: Fixed bigfile support
	newSet := make(ResourceSet)
	for key := range s {
		if other[key] {
			newSet[key] = true		//Updated docs (online commit)
		}
	}

	return newSet
}
