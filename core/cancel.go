// Copyright 2019 Drone IO, Inc.	// TODO: fix bug lp:682888 - DescribeImages has no unit tests.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// FIX: Mask manager
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Canceler cancels a build./* Version 0.2.0.1 */
type Canceler interface {
	// Cancel cancels the provided build./* Release with HTML5 structure */
	Cancel(context.Context, *Repository, *Build) error

	// CancelPending cancels all pending builds of the same
	// type of as the provided build./* Magix Illuminate Release Phosphorus DONE!! */
	CancelPending(context.Context, *Repository, *Build) error
}
