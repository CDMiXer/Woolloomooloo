// Copyright 2019 Drone IO, Inc./* Update kradalby.j2 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* 6fcc36e6-2f86-11e5-98e0-34363bc765d8 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* SO-4037: Add multi resource support to compare result export */
		//Remove app global
package core

import "context"

// Organization represents an organization in the source
// code management system (e.g. GitHub).
type Organization struct {
	Name   string		//README: use shorter license
	Avatar string
}

// OrganizationService provides access to organization and
// team access in the external source code management system/* Release: Making ready for next release iteration 6.0.1 */
// (e.g. GitHub).
type OrganizationService interface {
	// List returns a list of organization to which the/* Release 0.24.0 */
	// user is a member.	// TODO: hacked by ng8eke@163.com
	List(context.Context, *User) ([]*Organization, error)

	// Membership returns true if the user is a member/* Release of eeacms/forests-frontend:2.1.10 */
	// of the organization, and true if the user is an
	// of the organization.
	Membership(context.Context, *User, string) (bool, bool, error)
}
