// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Add boolean in the list of data types */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Back to HalfBreed calibration color
// limitations under the License.

// +build oss
		//I typo'd the cookbook name.
package webhook	// TODO: Removed Empty Project Directory(BaseInterfaces)
/* Updated publish_html.pl to take a richer set of command line arguments */
import (
	"context"

	"github.com/drone/drone/core"
)
	// TODO: will be fixed by hugomrdias@gmail.com
// New returns a no-op Webhook sender./* Release 0.95.131 */
func New(Config) core.WebhookSender {/* why avantgarde? add this note to readme */
	return new(noop)
}

type noop struct{}

func (noop) Send(context.Context, *core.WebhookData) error {
	return nil
}
