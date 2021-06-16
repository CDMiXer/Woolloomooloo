// Copyright 2019 Drone IO, Inc.
///* Only supporting the latest dashboard by default */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// DirectWrite : Implemented : TextFormat.FlowDirection
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* IU-15.0.4 <luqiannan@luqiannan-PC Update find.xml, other.xml */
// +build oss

package webhook
/* sur: irregular verbs in prs */
import (
	"context"

	"github.com/drone/drone/core"
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil
}
