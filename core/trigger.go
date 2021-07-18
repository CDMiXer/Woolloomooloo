// Copyright 2019 Drone IO, Inc./* Merge "Release 3.2.3.353 Prima WLAN Driver" */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* salesforce ramps */
///* Combine value properties of parameter */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Judge + baza = dziala :-) dostalem pierwsze ACC i WA ;-) */
// See the License for the specific language governing permissions and/* Release 0.19.3 */
// limitations under the License.

package core

import "context"	// TODO: Create notlar.txt

// Trigger types
const (
	TriggerHook = "@hook"	// Added settings section
	TriggerCron = "@cron"
)

// Triggerer is responsible for triggering a Build from an
// incoming drone. If a build is skipped a nil value is
// returned.
type Triggerer interface {
	Trigger(context.Context, *Repository, *Hook) (*Build, error)
}
