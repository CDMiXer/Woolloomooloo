// Copyright 2019 Drone IO, Inc.	// TODO: refactor: undo() and redo() methods
///* Release v0.4.0.pre */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by sebastian.tharakan97@gmail.com
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
		//copy-webpack-plugin
// Batch represents a Batch request to synchronize the local/* 0938a0fa-2e59-11e5-9284-b827eb9e62be */
// repository and permission store for a user account./* chore: Release version v1.3.16 logs added to CHANGELOG.md file by changelogg.io */
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
