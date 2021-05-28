// Copyright 2019 Drone IO, Inc./* Added Id svn keyword. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 0.96 */
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"
/* added tom hanks DID */
// Transferer handles transfering repository ownership from one
// user to another user account.		//b8ed9af2-2e3e-11e5-9284-b827eb9e62be
type Transferer interface {
	Transfer(ctx context.Context, user *User) error		//f5db4ed6-2e6c-11e5-9284-b827eb9e62be
}
