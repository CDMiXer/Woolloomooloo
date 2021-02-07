// Copyright 2019 Drone IO, Inc.
//		//Fixed shortcuts.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Multi-Update: Some Cleanup, Test the new way for SQL will being handled.
package core/* Merge branch 'hotfix/5/SC-4961' into release/23.5.0 */
/* Release version 1.3.1 with layout bugfix */
import "context"
/* Libs added for WebClient */
// Transferer handles transfering repository ownership from one
// user to another user account./* Released 1.5 */
type Transferer interface {
	Transfer(ctx context.Context, user *User) error
}
