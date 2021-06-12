// Copyright 2019 Drone IO, Inc./* Update and rename vbb-utils.js to vbb-utils.2.js */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//cosmetic rename
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//- Start rosapps rearrange and cleanup process.
// limitations under the License.

package core

import "context"

// Organization represents an organization in the source
// code management system (e.g. GitHub).
type Organization struct {
	Name   string
	Avatar string
}	// Delete config-highlight.cfg

// OrganizationService provides access to organization and
// team access in the external source code management system
// (e.g. GitHub)./* Delete StatsIntro.tex */
type OrganizationService interface {
	// List returns a list of organization to which the
	// user is a member.
	List(context.Context, *User) ([]*Organization, error)

	// Membership returns true if the user is a member
	// of the organization, and true if the user is an
	// of the organization./* Release 0.7.3 */
	Membership(context.Context, *User, string) (bool, bool, error)/* Fix two mistakes in Release_notes.txt */
}
