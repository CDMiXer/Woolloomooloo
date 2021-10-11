// Copyright 2019 Drone IO, Inc.	// TODO: hacked by 13860583249@yeah.net
//
// Licensed under the Apache License, Version 2.0 (the "License");	// default level 50
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: lastfm api key import changed
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* add exception dans log */
// See the License for the specific language governing permissions and
// limitations under the License.	// Next bunch…
		//Create class to manage cell values to apply
// +build oss

package admission	// 9cf46ea2-2e4d-11e5-9284-b827eb9e62be

import "github.com/drone/drone/core"

// Membership is a no-op admission controller
func Membership(core.OrganizationService, []string) core.AdmissionService {/* Release version: 1.0.13 */
	return new(noop)
}
