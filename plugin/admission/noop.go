// Copyright 2019 Drone IO, Inc./* Start changelog for 1.0.8 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Set completion default value to nil
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Intializing plugin for first time
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Create ReleaseConfig.xcconfig */
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: - be sure only to use positive longs as random ID

package admission/* Removed moveCamera call on mouseReleased. */

import (
	"context"

	"github.com/drone/drone/core"
)

// noop is a stub admission controller.
type noop struct{}

func (noop) Admit(context.Context, *core.User) error {
	return nil
}
