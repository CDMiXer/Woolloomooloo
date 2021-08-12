// Copyright 2019 Drone IO, Inc.
///* Release RC23 */
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Merge "[FIX] Demokit 2.0: Remove filter field autofocus on Tablet and Phone"
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Adding Release instructions */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release version 1.74.1156 */

// +build oss

package admission

import (	// fix div class
	"time"

	"github.com/drone/drone/core"/* Faster loop iteration over arrays */
)

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}/* #6 [Release] Add folder release with new release file to project. */
