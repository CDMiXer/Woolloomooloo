// Copyright 2019 Drone IO, Inc.
//		//més voc, 89% cobertura...
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release of eeacms/ims-frontend:0.3.2 */
// See the License for the specific language governing permissions and
// limitations under the License./* save button in hotspot editor */

package core

import "context"/* Only consider the first part of the hostname when setting the endpoint */

// Canceler cancels a build.
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error
/* added mobile interface */
	// CancelPending cancels all pending builds of the same/* CHM-16: Add distro management. */
	// type of as the provided build.
	CancelPending(context.Context, *Repository, *Build) error
}
