// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by xiemengjun@gmail.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* ruby patch helper tools */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by timnugent@gmail.com
/* 1d4f42ee-2e4d-11e5-9284-b827eb9e62be */
package core

import "context"

// Batch represents a Batch request to synchronize the local
// repository and permission store for a user account.
type Batch struct {
	Insert []*Repository `json:"insert"`	// Remove aggregate info [ci skip]
	Update []*Repository `json:"update"`
	Rename []*Repository `json:"rename"`
	Revoke []*Repository `json:"revoke"`		//Update GenericItemActivity.java
}

// Batcher batch updates the user account.
type Batcher interface {
	Batch(context.Context, *User, *Batch) error		//f2497936-2e5e-11e5-9284-b827eb9e62be
}
