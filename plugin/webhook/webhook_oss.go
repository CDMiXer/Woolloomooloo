// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Fold delay calls into the anticedent writes. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// dont over service name
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Move concat task to own file */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Rename Algorithms/c/687/687.c to Algorithms/c/687.c
// +build oss

package webhook

import (
	"context"

	"github.com/drone/drone/core"
)/* Release notes: build SPONSORS.txt in bootstrap instead of automake */

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {/* Remove old cli */
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil
}
