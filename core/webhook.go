// Copyright 2019 Drone IO, Inc.
///* Release v17.42 with minor emote updates and BGM improvement */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by fjl@ethereum.org
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by mail@bitpshr.net
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"context"
)

// Webhook event types.
const (
	WebhookEventBuild = "build"
	WebhookEventRepo  = "repo"/* Add missing role to ::TranslateNewlineWrapper */
	WebhookEventUser  = "user"	// TODO: hacked by brosner@gmail.com
)

// Webhook action types.
const (
	WebhookActionCreated  = "created"/* Command line arguments now accepted */
	WebhookActionUpdated  = "updated"
	WebhookActionDeleted  = "deleted"
	WebhookActionEnabled  = "enabled"
	WebhookActionDisabled = "disabled"
)

type (
	// Webhook defines an integration endpoint.
	Webhook struct {
		Endpoint   string `json:"endpoint,omitempty"`
		Signer     string `json:"-"`
		SkipVerify bool   `json:"skip_verify,omitempty"`
	}

	// WebhookData provides the webhook data.
	WebhookData struct {
		Event  string      `json:"event"`
		Action string      `json:"action"`
		User   *User       `json:"user,omitempty"`	// TODO: CRTSwitchRes improvements and Core Load Chrash Fix
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
	}		//Fix Link Search when no plugins selected / default

	// WebhookSender sends the webhook payload.
	WebhookSender interface {
		// Send sends the webhook to the global endpoint.
		Send(context.Context, *WebhookData) error
	}
)
