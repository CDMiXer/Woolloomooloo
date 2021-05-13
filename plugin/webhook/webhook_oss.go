// Copyright 2019 Drone IO, Inc./* Merge "Release 3.2.3.374 Prima WLAN Driver" */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Removed obsolete include of <boost/concept_check.hpp>. */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Updated the version number. Added an option in the scoring preferences.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// Mobile UI WIP

// +build oss

package webhook
	// TODO: fixup task
import (
	"context"

	"github.com/drone/drone/core"
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {	// TODO: hacked by alan.shaw@protocol.ai
	return new(noop)/* Merge "Update Release notes for 0.31.0" */
}/* Release 1.2.11 */

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil
}
