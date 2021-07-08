// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Delete IpfCcmBoPgLoElementUpdateResponse.java */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//DocExtract: LHCHXSWG-INT report numbers
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Create p81-p84.lisp
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Trigger types	// TODO: Put V5/4 first, V3 later
const (	// TODO: Delete cs_logger and associated code
	TriggerHook = "@hook"
	TriggerCron = "@cron"
)

// Triggerer is responsible for triggering a Build from an
// incoming drone. If a build is skipped a nil value is
// returned.
type Triggerer interface {
	Trigger(context.Context, *Repository, *Hook) (*Build, error)
}
