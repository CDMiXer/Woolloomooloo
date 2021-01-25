// Copyright 2019 Drone IO, Inc.		//Moved defaulting of args.device into ArgumentParser set up.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Fix the lack of newline information
// You may obtain a copy of the License at
///* Removed bower dependency for angular-bootstrap */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Should fix "Vault-api" issue.
// distributed under the License is distributed on an "AS IS" BASIS,/* 43505a18-2e67-11e5-9284-b827eb9e62be */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss		//CWS-TOOLING: integrate CWS jl146_DEV300

package webhook

import (
	"context"

	"github.com/drone/drone/core"
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)/* Merge "Migrate to stevedore" */
}

type noop struct{}
	// TODO: will be fixed by peterke@gmail.com
func (noop) Send(context.Context, *core.WebhookData) error {
	return nil
}
