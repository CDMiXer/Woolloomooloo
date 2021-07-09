// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// remove merge confilct
// You may obtain a copy of the License at/* Pull up common XML feature methods */
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Released v1.3.5 */
//
// Unless required by applicable law or agreed to in writing, software/* Fixed return type of wait for js and wait for jquery condition keywords */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//[IMP]remove repeated code
// +build oss

package admission

import (
	"time"

	"github.com/drone/drone/core"
)		//Update DeltaSyncRunner.java

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}
