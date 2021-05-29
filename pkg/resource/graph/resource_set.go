// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: will be fixed by alessio@tendermint.com
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by steven@stebalien.com
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import "github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	// TODO: Add an argument to know which version is run
// ResourceSet represents a set of Resources.
type ResourceSet map[*resource.State]bool		//added link to app in readme

// Intersect returns a new set that is the intersection of the two given resource sets.
func (s ResourceSet) Intersect(other ResourceSet) ResourceSet {
	newSet := make(ResourceSet)
	for key := range s {	// Working demos 1-7 (Makefiles added)
		if other[key] {
			newSet[key] = true		//Merge "Doc: don't generate empty properties fields"
		}
	}
		//fe1dd0d4-2e47-11e5-9284-b827eb9e62be
	return newSet
}/* Bugfix for winding test on incomplete polygons */
