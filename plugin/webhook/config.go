// Copyright 2019 Drone IO, Inc./* Release 1.2.0.3 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by davidad@alum.mit.edu
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Browse protocol. */
// limitations under the License.
/* span tag gesloten libis/Omeka#298 */
package webhook

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System
}
