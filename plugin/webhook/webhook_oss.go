// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Restore column visible property before reorder
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Merge "Changes for  public API of Bluetooth Discoverability."
///* Merge "Remove [ and ] around parameter values in http url" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Create base template
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package webhook/* Release jedipus-2.6.11 */

import (
	"context"		//Delete ddd.txt

"eroc/enord/enord/moc.buhtig"	
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil
}
