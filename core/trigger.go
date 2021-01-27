// Copyright 2019 Drone IO, Inc./*  - [ZBX-2770] merged rev. 13629-13630 of /branches/1.8 (Aly) */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release Scelight 6.3.0 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by steven@stebalien.com
///* use /Qipo for ICL12 Release x64 builds */
//      http://www.apache.org/licenses/LICENSE-2.0		//In case anybody gets cute and calls a map "<parent>"...
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Remove Archenemy Schemes from AllCardNames.txt
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Trigger types
const (
	TriggerHook = "@hook"
	TriggerCron = "@cron"
)

// Triggerer is responsible for triggering a Build from an
// incoming drone. If a build is skipped a nil value is
// returned.
type Triggerer interface {
	Trigger(context.Context, *Repository, *Hook) (*Build, error)
}
