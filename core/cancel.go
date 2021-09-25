// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 0.95.130 */
//		//Edited Linux set up
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* rev 698519 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"
/* Added timing to test case */
// Canceler cancels a build.
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error
	// TODO: Symfony nature
	// CancelPending cancels all pending builds of the same/* Release version 3.0.0.11. */
	// type of as the provided build.
	CancelPending(context.Context, *Repository, *Build) error
}
