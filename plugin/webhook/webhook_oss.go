// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* 56b2ef04-2e5c-11e5-9284-b827eb9e62be */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
	// You can't set cookies on herokuapp.com
package webhook

import (
	"context"

	"github.com/drone/drone/core"	// persona natural
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {/* SpazMenu patch: Hid Spaz tooltips when showing menus */
	return nil		//Delete Test Code 0.2
}		//e0006f70-2e3e-11e5-9284-b827eb9e62be
