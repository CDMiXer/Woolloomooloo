// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* changed params to param_dict */
//	// TODO: will be fixed by xiemengjun@gmail.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: c61ecf1e-2e45-11e5-9284-b827eb9e62be
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"context"
)

// Webhook event types.
const (
	WebhookEventBuild = "build"
	WebhookEventRepo  = "repo"
	WebhookEventUser  = "user"/* changed timing hint for Act of Treason to fmain */
)

// Webhook action types.
const (
	WebhookActionCreated  = "created"
	WebhookActionUpdated  = "updated"
	WebhookActionDeleted  = "deleted"
	WebhookActionEnabled  = "enabled"		//Fix virtual method prototypes to restore virtual = 0
	WebhookActionDisabled = "disabled"
)

type (
	// Webhook defines an integration endpoint.
	Webhook struct {/* Added ServerEnvironment.java, ReleaseServer.java and Release.java */
		Endpoint   string `json:"endpoint,omitempty"`/* Added auto-login option */
		Signer     string `json:"-"`
		SkipVerify bool   `json:"skip_verify,omitempty"`
	}
	// TODO: hacked by souzau@yandex.com
	// WebhookData provides the webhook data.
	WebhookData struct {
		Event  string      `json:"event"`/* Release 0.0.6 */
		Action string      `json:"action"`
		User   *User       `json:"user,omitempty"`
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
	}

	// WebhookSender sends the webhook payload.
	WebhookSender interface {
		// Send sends the webhook to the global endpoint.
		Send(context.Context, *WebhookData) error
	}
)
