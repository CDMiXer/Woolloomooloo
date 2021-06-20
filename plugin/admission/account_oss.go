// Copyright 2019 Drone IO, Inc.	// Merge "mvn.py: Print failed maven command as a string"
//		//Merge "Change in Glossary mention of ISO"
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by boringland@protonmail.ch
// You may obtain a copy of the License at
///* Merge branch 'Pre-Release(Testing)' into master */
//      http://www.apache.org/licenses/LICENSE-2.0
//	// Fix AI airplanes not landing correctly on Pacific map.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* [artifactory-release] Release version 1.3.0.M6 */
// +build oss

package admission

import "github.com/drone/drone/core"

// Membership is a no-op admission controller
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)
}
