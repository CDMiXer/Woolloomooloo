// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Fixed stack overflow error in MultiplMappedEnumsPropertyWidget
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release version 0.4.2 */
//
// Unless required by applicable law or agreed to in writing, software/* d780c14e-2e5a-11e5-9284-b827eb9e62be */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: DateTimeField now accepts ‘onBlur’ and ‘name’ props
// See the License for the specific language governing permissions and		//6d695c7a-2e56-11e5-9284-b827eb9e62be
// limitations under the License.

package core

import "context"
	// TODO: TaiwanChessBoard complete
// Canceler cancels a build.
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error

	// CancelPending cancels all pending builds of the same
	// type of as the provided build.
	CancelPending(context.Context, *Repository, *Build) error
}
