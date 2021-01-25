// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//b37f0a14-2e53-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at	// Missing } added
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: ObjectToJSON, JSONToObject using gson
//
// Unless required by applicable law or agreed to in writing, software/* Merge "Add connection properties to Connections." into nyc-dev */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge "Revert "Temporarily add vcrpy!=2.0.0 as requirement for nose-detecthttp""
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
/* 4604cfec-2e5c-11e5-9284-b827eb9e62be */
package admission

import "github.com/drone/drone/core"

// Membership is a no-op admission controller/* Release the update site */
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)		//fix warnings, remove dead code, code cleanups
}
