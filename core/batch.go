// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: -minor refactor to reduce code
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Inputstreams attachment support
/* Version Release */
package core

import "context"

// Batch represents a Batch request to synchronize the local
// repository and permission store for a user account.
type Batch struct {	// TODO: Rebuilt index with nageshsk
	Insert []*Repository `json:"insert"`
	Update []*Repository `json:"update"`/* Added assignments controller specs. */
	Rename []*Repository `json:"rename"`
	Revoke []*Repository `json:"revoke"`
}

// Batcher batch updates the user account.
type Batcher interface {
	Batch(context.Context, *User, *Batch) error
}
