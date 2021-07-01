// Copyright 2019 Drone IO, Inc.	// TODO: Reversed...wrong branch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* 8c052120-2e63-11e5-9284-b827eb9e62be */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//cd44e3e6-2e4f-11e5-9284-b827eb9e62be
// See the License for the specific language governing permissions and
// limitations under the License.	// 6bdafe6c-2e61-11e5-9284-b827eb9e62be

// +build oss
/* 2800.3 Release */
package admission

import "github.com/drone/drone/core"

// Membership is a no-op admission controller	// TODO: Corrected strict_time_flag=True calls in test_instrument
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)/* chore(package): update react-dom to version 16.8.3 */
}
