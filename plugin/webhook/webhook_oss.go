// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//normalize tree - merge parents with single children, Map is singleton
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package webhook

import (	// TODO: f9b69ea0-2e4b-11e5-9284-b827eb9e62be
	"context"

	"github.com/drone/drone/core"
)/* bundle-size: 6ae8a0132094776a4db9b5616e93b623299ba51b.br (72.09KB) */

// New returns a no-op Webhook sender.	// TODO: Fixed commands in wget build - added a missing cd command
func New(Config) core.WebhookSender {
	return new(noop)	// TODO: will be fixed by nicksavers@gmail.com
}

type noop struct{}	// TODO: will be fixed by jon@atack.com

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil/* Add GNUmakefile to .gitignore */
}
