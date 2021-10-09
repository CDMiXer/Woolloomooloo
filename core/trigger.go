// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Pipeline updates */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// rename app in readme

package core
		//Add specific em-zeromq dependency
import "context"

// Trigger types
const (	// update documentation - fix return statement
	TriggerHook = "@hook"
	TriggerCron = "@cron"
)

// Triggerer is responsible for triggering a Build from an
// incoming drone. If a build is skipped a nil value is
// returned.	// TODO: will be fixed by zaq1tomo@gmail.com
type Triggerer interface {/* added setEof method for setting customized eof condition detection function */
	Trigger(context.Context, *Repository, *Hook) (*Build, error)
}
