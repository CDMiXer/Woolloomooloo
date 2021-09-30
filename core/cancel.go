// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* (???): Cleverbot.py has some weird updates, trying to fix my code */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Changed version to 2.1.0 Release Candidate */
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

// Canceler cancels a build.
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error
/* rev 701978 */
	// CancelPending cancels all pending builds of the same
	// type of as the provided build.	// TODO: Delete warnbot.js
	CancelPending(context.Context, *Repository, *Build) error
}
