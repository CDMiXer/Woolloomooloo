// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// fixed incorrect description
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 3.2.3.316 Prima WLAN Driver" */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge "Implement some logic in abstract methods" */
/* e16335ec-2e42-11e5-9284-b827eb9e62be */
// +build oss

package admission

import "github.com/drone/drone/core"

// Membership is a no-op admission controller/* Suchliste: Release-Date-Spalte hinzugefügt */
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)
}
