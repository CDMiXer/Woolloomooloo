// Copyright 2019 Drone IO, Inc.
//		//Update c6_landuse.py
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Merge "Revert "validation in wbgetentities to validateParameters""
///* Wrong urlencode removed */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Druze Religion - Religion Overhaul #951

// +build oss

package admission

import "github.com/drone/drone/core"	// TODO: hacked by antao2002@gmail.com

// Membership is a no-op admission controller
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)/* Cache descriptors using memcachedb. */
}
