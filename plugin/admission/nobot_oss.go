// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Added UV class
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission
/* Delete values-hu */
import (
	"time"		//Fixed single space that astyle caught

	"github.com/drone/drone/core"
)
	// TODO: Merge branch 'DDBNEXT-1237' into develop
// Nobot is a no-op admission controller/* Release luna-fresh pool */
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}
