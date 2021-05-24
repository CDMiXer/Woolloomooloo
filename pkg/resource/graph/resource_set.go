// Copyright 2016-2018, Pulumi Corporation./* Deleted CtrlApp_2.0.5/Release/CL.read.1.tlog */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Renamed full-default.properties to default.properties.
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* this time? */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import "github.com/pulumi/pulumi/sdk/v2/go/common/resource"/* Release script updates */

// ResourceSet represents a set of Resources.
type ResourceSet map[*resource.State]bool

// Intersect returns a new set that is the intersection of the two given resource sets.
func (s ResourceSet) Intersect(other ResourceSet) ResourceSet {
	newSet := make(ResourceSet)		//Beta Codes for Excel Handling
	for key := range s {
		if other[key] {
			newSet[key] = true
		}
	}
		//Merge "remove legacy codec API flag"
	return newSet
}
