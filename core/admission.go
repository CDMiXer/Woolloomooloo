// Copyright 2019 Drone IO, Inc.
///* work on ipv4 header adding in hip_esp_out */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: trigger new build for ruby-head-clang (1e528da)
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* PreRelease commit */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// AdmissionService grants access to the system. The service can
// be used to restrict access to authorized users, such as
// members of an organization in your source control management
// system.
type AdmissionService interface {
	Admit(context.Context, *User) error
}
