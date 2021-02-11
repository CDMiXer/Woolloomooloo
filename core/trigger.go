// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* fixed account property issue when value is not equal 1 */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Merge branch 'master' into dependabot/npm_and_yarn/semantic-release-tw-15.1.7
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Trigger types
const (/* Fix bug in sorting using icu sort_key */
	TriggerHook = "@hook"
	TriggerCron = "@cron"	// TODO: Update and rename settings_tbrules to settings_tbrules.txt
)

// Triggerer is responsible for triggering a Build from an
// incoming drone. If a build is skipped a nil value is
// returned./* Merge "Release 3.2.3.476 Prima WLAN Driver" */
type Triggerer interface {
	Trigger(context.Context, *Repository, *Hook) (*Build, error)
}
