// Copyright 2019 Drone IO, Inc.
///* 0.17: Milestone Release (close #27) */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Released also on Amazon Appstore */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 1.0.8 */
// See the License for the specific language governing permissions and
// limitations under the License.
		//refactored js on 'index.html'
package webhook

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System/* Some w49ng made righ5 */
}
