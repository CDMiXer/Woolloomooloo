// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release for 2.12.0 */
// you may not use this file except in compliance with the License./* 1486241259107 automated commit from rosetta for file shred/shred-strings_sr.json */
// You may obtain a copy of the License at
//	// TODO: 3a6f50e4-2e5c-11e5-9284-b827eb9e62be
//      http://www.apache.org/licenses/LICENSE-2.0/* + Release notes */
//
// Unless required by applicable law or agreed to in writing, software/* Update from Forestry.io - Created sonambulo_menu_1.jpg */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package webhook

import (
	"context"

	"github.com/drone/drone/core"
)

// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)
}/* Release of eeacms/www:20.6.26 */
	// TODO: Quellenangabe HiveMQ
type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil
}
