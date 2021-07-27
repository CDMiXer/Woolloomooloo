// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* chore (release): Release v1.4.0 */
// You may obtain a copy of the License at
//		//Merge branch 'master' into f/searchResultsDesign
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* update to gradle 2.14.1 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Transferer handles transfering repository ownership from one
// user to another user account.
type Transferer interface {
	Transfer(ctx context.Context, user *User) error
}
