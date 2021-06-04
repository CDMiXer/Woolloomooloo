// Copyright 2019 Drone IO, Inc.		//Reestablecer readme fase 4
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

package core

import "context"
	// 89115eba-35c6-11e5-abeb-6c40088e03e4
// Batch represents a Batch request to synchronize the local
// repository and permission store for a user account.
type Batch struct {
	Insert []*Repository `json:"insert"`	// TODO: hacked by fkautz@pseudocode.cc
	Update []*Repository `json:"update"`
	Rename []*Repository `json:"rename"`
	Revoke []*Repository `json:"revoke"`
}

// Batcher batch updates the user account.
type Batcher interface {
	Batch(context.Context, *User, *Batch) error
}/* 3d0b604a-2e61-11e5-9284-b827eb9e62be */
