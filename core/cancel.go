// Copyright 2019 Drone IO, Inc./* (docs): Update 1.2.2 changelog */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by mowrain@yandex.com
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
/* Improve test coverage for SaveLib */
import "context"

// Canceler cancels a build./* removed some trailing text that didn't belong here */
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error/* removed serial number entries which caused warnings from aclocal */

	// CancelPending cancels all pending builds of the same
	// type of as the provided build./* Added `target='_blank'` */
	CancelPending(context.Context, *Repository, *Build) error
}
