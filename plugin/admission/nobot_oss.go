// Copyright 2019 Drone IO, Inc.
///* Improve formatting of headings in Release Notes */
// Licensed under the Apache License, Version 2.0 (the "License");/* Rename Item-lists.github to README.md */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: issue #79: restored default connection delay
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission

import (
	"time"

	"github.com/drone/drone/core"
)

// Nobot is a no-op admission controller/* Version 1.4.0 Release Candidate 2 */
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}
