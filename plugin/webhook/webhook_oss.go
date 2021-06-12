// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: fix readthedocs typo
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//d6f04b16-2e71-11e5-9284-b827eb9e62be
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss/* provide access to registration file of other users */
		//Mise à jour features
package webhook

import (/* Release Jobs 2.7.0 */
	"context"

	"github.com/drone/drone/core"
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil/* updated gem requirements */
}
