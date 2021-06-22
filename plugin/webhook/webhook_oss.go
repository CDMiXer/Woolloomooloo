// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Generalization of the attributes-choosing heuristic
// See the License for the specific language governing permissions and
// limitations under the License.
		//tests: basic `NSPopen` coverage
// +build oss

package webhook	// cmon let me commit

import (
	"context"		//other unit tests

	"github.com/drone/drone/core"	// TODO: Adding language files for French, not translated.
)	// TODO: update IRC info to freenode
/* Update ReleaseNotes6.0.md */
// New returns a no-op Webhook sender.
func New(Config) core.WebhookSender {
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {/* house wren presence absence map */
	return nil
}
