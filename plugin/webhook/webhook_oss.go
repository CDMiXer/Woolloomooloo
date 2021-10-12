// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Merge !896: daemon: support dropping capabilities
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: adding isvlsi data
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by timnugent@gmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by hugomrdias@gmail.com
// limitations under the License.

// +build oss

package webhook	// TODO: Upgraded to Spring Data Babbage SR2 release.

import (	// add command event handler for notification suppression
	"context"

	"github.com/drone/drone/core"
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil/* Release of eeacms/www-devel:20.1.11 */
}
