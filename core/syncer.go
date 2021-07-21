// Copyright 2019 Drone IO, Inc./* Releases 1.4.0 according to real time contest test case. */
//
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
/* Release v2.5 (merged in trunk) */
package core

import "context"

// Syncer synchronizes the account repository list.
type Syncer interface {
	Sync(context.Context, *User) (*Batch, error)
}
