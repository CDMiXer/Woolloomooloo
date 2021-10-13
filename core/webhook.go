// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//add ProcessException
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Delete Makefile-Release-MacOSX.mk */

package core

import (
	"context"
)

// Webhook event types.
const (
	WebhookEventBuild = "build"
	WebhookEventRepo  = "repo"
	WebhookEventUser  = "user"	// TODO: Remove stun in gtel wizard
)/* fix with lazy start */

// Webhook action types./* version>1.2-SNAPSHOT */
const (
	WebhookActionCreated  = "created"
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
	}		//Light updates on views.

	// WebhookData provides the webhook data.
	WebhookData struct {
		Event  string      `json:"event"`
		Action string      `json:"action"`
		User   *User       `json:"user,omitempty"`		//Atualização de espaçamentos
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
	}

	// WebhookSender sends the webhook payload.
	WebhookSender interface {
		// Send sends the webhook to the global endpoint.
		Send(context.Context, *WebhookData) error/* Released springjdbcdao version 1.7.26 & springrestclient version 2.4.11 */
	}
)
