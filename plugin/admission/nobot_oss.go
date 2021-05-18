// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: README: Use fenced code blocks with syntax highlighting
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//fixed loop again - check is on ForwardSolver now!
// limitations under the License.	// TODO: Merge "Pulling out predictions into another row view." into ub-launcher3-burnaby

// +build oss

package admission

import (/* Fixed functions' name in oscam.h/oscam.c */
	"time"

	"github.com/drone/drone/core"
)

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}
