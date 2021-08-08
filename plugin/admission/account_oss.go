// Copyright 2019 Drone IO, Inc.
//		//Create AdnForme11.cpp
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by martin2cai@hotmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: Fix broken de/serialization for a couple of C++ Exprs.
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* a091e5ae-2e65-11e5-9284-b827eb9e62be */
// limitations under the License.

// +build oss

package admission

import "github.com/drone/drone/core"	// TODO: Update Whoaverse.css
	// Update aplicaciones.sh
// Membership is a no-op admission controller
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)
}
