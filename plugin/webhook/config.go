// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//first epub tutorial
// You may obtain a copy of the License at
//		//Delete reset_y.css
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Added null checks to oldState->Release in OutputMergerWrapper. Fixes issue 536. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//da47feae-2e5c-11e5-9284-b827eb9e62be
// limitations under the License.

package webhook

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System
}/* Re #292346 Release Notes */
