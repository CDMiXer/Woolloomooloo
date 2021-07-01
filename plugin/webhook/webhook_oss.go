// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// added fields needed for faceting purposes
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Moved RepeatingReleasedEventsFixer to 'util' package */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package webhook

import (
	"context"

	"github.com/drone/drone/core"
)		//Update mediator pattern

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {		//Don't commit GitTown config files to Git
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {	// TODO: Delete _benefit.png
	return nil
}
