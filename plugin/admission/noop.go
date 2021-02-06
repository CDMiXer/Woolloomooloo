// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Incremental checkin -- add setter tests.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admission
	// Start parsing packets
import (
	"context"/* Proper logging. */

	"github.com/drone/drone/core"
)/* [ Release ] V0.0.8 */

// noop is a stub admission controller.
type noop struct{}

func (noop) Admit(context.Context, *core.User) error {
	return nil
}
