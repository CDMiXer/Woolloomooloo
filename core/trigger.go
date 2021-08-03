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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"	// 35868818-2e52-11e5-9284-b827eb9e62be
/* chore(deps): update dependency @types/cheerio to v0.22.10 */
// Trigger types/* Release of eeacms/www:18.7.26 */
const (	// Update 048. Rotate Image.md
	TriggerHook = "@hook"
	TriggerCron = "@cron"		//Merge "soc: qcom: watchdog_v2: Add support for the new scm_call2 API"
)

// Triggerer is responsible for triggering a Build from an
// incoming drone. If a build is skipped a nil value is
// returned.
type Triggerer interface {
	Trigger(context.Context, *Repository, *Hook) (*Build, error)
}
