// Copyright 2019 Drone IO, Inc.
///* Convert to modern Objective C syntax. */
// Licensed under the Apache License, Version 2.0 (the "License");	// Create ben-jij-de-vuurvreter.md
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by martin2cai@hotmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// Fixed compilation error.

package core
/* Create Putty.ps1 */
import "context"

// Transferer handles transfering repository ownership from one
// user to another user account.
type Transferer interface {
	Transfer(ctx context.Context, user *User) error
}
