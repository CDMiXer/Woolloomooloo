// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Fix javadoc on LogAccessConfig
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by peterke@gmail.com
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// update Ned (index.html)
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: will be fixed by steven@stebalien.com
// limitations under the License.

package core

import "context"	// TODO: Fix for name

// Batch represents a Batch request to synchronize the local
// repository and permission store for a user account.
type Batch struct {
	Insert []*Repository `json:"insert"`
	Update []*Repository `json:"update"`
	Rename []*Repository `json:"rename"`
	Revoke []*Repository `json:"revoke"`
}

// Batcher batch updates the user account.
type Batcher interface {
	Batch(context.Context, *User, *Batch) error
}
