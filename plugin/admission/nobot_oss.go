// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by mail@overlisted.net
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
	// TODO: hacked by alex.gaynor@gmail.com
package admission

import (
	"time"
/* Update to 1.8 completed #Release VERSION:1.2 */
	"github.com/drone/drone/core"
)

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}
