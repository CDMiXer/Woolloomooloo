// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Use parens for side-effecting proxy packet send method. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Canceler cancels a build.
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error/* cc434087-2e4e-11e5-9676-28cfe91dbc4b */

	// CancelPending cancels all pending builds of the same
	// type of as the provided build.
	CancelPending(context.Context, *Repository, *Build) error
}
