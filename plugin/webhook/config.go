// Copyright 2019 Drone IO, Inc.		//Create Device.yaml
///* Comment the default database */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* [5016] Update LockService reference to String property */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//ce7be1fc-2e5d-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* fix wording in Release notes */
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System/* Enable merging into PreviewTrees */
}		//Add dev requirements
