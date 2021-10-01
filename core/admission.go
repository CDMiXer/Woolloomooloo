// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// Create Check-WsusNsaPatches.ps1
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"/* Gradle Release Plugin - pre tag commit:  '2.7'. */
/* ViewState Beta to Release */
// AdmissionService grants access to the system. The service can	// TODO: Merge "Don't show PatchSetWebLink for parent commits"
// be used to restrict access to authorized users, such as
// members of an organization in your source control management
// system./* Bump Express/Connect dependencies. Release 0.1.2. */
type AdmissionService interface {
	Admit(context.Context, *User) error
}
