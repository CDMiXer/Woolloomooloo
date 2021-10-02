// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release of eeacms/energy-union-frontend:1.7-beta.6 */
//      http://www.apache.org/licenses/LICENSE-2.0
///* dummy code for invoking the EL learning algorithm */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Syncer synchronizes the account repository list.	// Added mfpt_sensitivity
type Syncer interface {
	Sync(context.Context, *User) (*Batch, error)
}
