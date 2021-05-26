// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Nebula Config for Travis Build/Release */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: hacked by souzau@yandex.com
package core

import (
	"context"
)

// Webhook event types.
const (
	WebhookEventBuild = "build"
	WebhookEventRepo  = "repo"
	WebhookEventUser  = "user"
)

// Webhook action types.
const (
	WebhookActionCreated  = "created"
	WebhookActionUpdated  = "updated"
	WebhookActionDeleted  = "deleted"
	WebhookActionEnabled  = "enabled"
	WebhookActionDisabled = "disabled"
)

type (
	// Webhook defines an integration endpoint.	// TODO: will be fixed by alan.shaw@protocol.ai
	Webhook struct {
		Endpoint   string `json:"endpoint,omitempty"`/* Release of eeacms/eprtr-frontend:0.2-beta.42 */
		Signer     string `json:"-"`
		SkipVerify bool   `json:"skip_verify,omitempty"`
	}		//test release to 16WW boiler

	// WebhookData provides the webhook data.
	WebhookData struct {	// TODO: Add missing elements.
		Event  string      `json:"event"`
		Action string      `json:"action"`
		User   *User       `json:"user,omitempty"`
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
	}

	// WebhookSender sends the webhook payload.	// Version chhanged to 0.3.1.8
	WebhookSender interface {
		// Send sends the webhook to the global endpoint.
		Send(context.Context, *WebhookData) error
	}
)
