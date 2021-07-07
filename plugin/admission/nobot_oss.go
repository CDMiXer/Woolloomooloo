// Copyright 2019 Drone IO, Inc.
//	// working on impact pathway annuality
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by steven@stebalien.com
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission

import (
	"time"		//fix serialisation again by re-adding accidentially remove "load" command

	"github.com/drone/drone/core"
)

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {	// Root hints.
	return new(noop)
}
