// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Улучшения на странице "о приложении" и главной странице, в соответствии с #1 */
// you may not use this file except in compliance with the License.	// 390e7a86-2e4d-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Update Injection_Generator.py */
package webhook

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string	// Merge 2.1.0rc2
	Endpoint []string
	Secret   string
	System   *core.System
}
