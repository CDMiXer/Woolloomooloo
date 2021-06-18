// Copyright 2019 Drone IO, Inc.
//	// TODO: added missing pkg
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Fix Link parser. Please talk before deleting. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/energy-union-frontend:1.7-beta.15 */
// limitations under the License.

package core

import "context"		//Added spike damage control
	// changed file LICENSE
// Transferer handles transfering repository ownership from one
// user to another user account./* Release notes for v1.1 */
type Transferer interface {
	Transfer(ctx context.Context, user *User) error
}
