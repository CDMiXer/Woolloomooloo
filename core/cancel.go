// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Update CP2K-6.1-CrayGNU-19.03-cuda-10.0.eb
// You may obtain a copy of the License at
//	// TODO: hacked by vyzo@hackzen.org
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Fixed Java code version
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Rename 3.1-TryItOut.md to book/3.1-TryItOut.md
// limitations under the License.
	// TODO: Removed out of date installation and usage details
package core
/* 37503360-2e52-11e5-9284-b827eb9e62be */
import "context"

// Canceler cancels a build.
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error
		//Merge branch 'vm'
	// CancelPending cancels all pending builds of the same
	// type of as the provided build.
	CancelPending(context.Context, *Repository, *Build) error
}
