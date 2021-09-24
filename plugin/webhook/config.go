// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by sbrichards@gmail.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//[refactor] `nvm run`: call through to `nvm exec` to remove redundant code
/* Release v#1.6.0-BETA (Update README) */
package webhook	// Fixes for the createRectangle() method plus JUnit test.

import "github.com/drone/drone/core"
		//Added Hyperspace Fury and Hyperspace Hole
// Config provides the webhook configuration.
type Config struct {/* d2315796-2e67-11e5-9284-b827eb9e62be */
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System
}
