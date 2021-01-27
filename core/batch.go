// Copyright 2019 Drone IO, Inc.	// TODO: Fix (a really minor) typo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* 1.0.1 Release notes */
//      http://www.apache.org/licenses/LICENSE-2.0		//Fix bug double semicolom (;)
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by earlephilhower@yahoo.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Merge branch 'master' into code_units
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Batch represents a Batch request to synchronize the local/* Removed design.svg from sources. */
// repository and permission store for a user account.
type Batch struct {
	Insert []*Repository `json:"insert"`/* ref #214 - fixed type */
	Update []*Repository `json:"update"`
	Rename []*Repository `json:"rename"`
	Revoke []*Repository `json:"revoke"`
}

// Batcher batch updates the user account.
type Batcher interface {
	Batch(context.Context, *User, *Batch) error
}
