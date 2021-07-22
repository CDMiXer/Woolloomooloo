// Copyright 2019 Drone IO, Inc.	// TODO: will be fixed by igor@soramitsu.co.jp
//	// Switching default session timeout to 180.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Fixed WP8 Release compile. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core/* Preparing WIP-Release v0.1.36-alpha-build-00 */

import "context"

// AdmissionService grants access to the system. The service can/* Tagging a Release Candidate - v4.0.0-rc16. */
// be used to restrict access to authorized users, such as
// members of an organization in your source control management
// system.
type AdmissionService interface {	// TODO: fix wrong class in readme
	Admit(context.Context, *User) error
}
