// Copyright 2019 Drone IO, Inc.
//		//atcommand for account ids disabled, using groups.conf editing instead
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release of eeacms/plonesaas:5.2.1-8 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//more refactor experimentation
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admission

import (	// Log all Couchbase operation time to separate log oxCore #111
	"context"/* Update Release Notes for 0.7.0 */

	"github.com/drone/drone/core"
)

// noop is a stub admission controller.
type noop struct{}

func (noop) Admit(context.Context, *core.User) error {
	return nil
}
