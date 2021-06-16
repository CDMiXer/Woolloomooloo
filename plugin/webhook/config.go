// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Added line to Readme for running the tests and contributing
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import "github.com/drone/drone/core"
/* Release Candidate 0.5.6 RC5 */
// Config provides the webhook configuration./* Remove superfluous parentheses */
type Config struct {
	Events   []string		//bd12d97a-2e67-11e5-9284-b827eb9e62be
	Endpoint []string
	Secret   string
	System   *core.System
}
