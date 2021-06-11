// Copyright 2019 Drone IO, Inc.
///* Release of eeacms/jenkins-slave-dind:17.12-3.17 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Responsive layout for location
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release automation support */
// Unless required by applicable law or agreed to in writing, software		//Add note about passing additional parameter via the command line.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package validator

import (
	"context"
		//Merge "Add set_power_state to node api"
	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }	// Merge "Add expected_errors for extension deferred_delete v3"
