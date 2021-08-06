// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Update custombootimg.mk
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//allow id tokens with no audience
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by steven@stebalien.com
// See the License for the specific language governing permissions and/* open Chrome URLs in Safari */
// limitations under the License./* Deposit Slip Editor updates */

package core

// Hook event constants.
const (
	EventCron        = "cron"	// TODO: hacked by boringland@protonmail.ch
	EventCustom      = "custom"
	EventPush        = "push"
	EventPullRequest = "pull_request"
	EventTag         = "tag"
	EventPromote     = "promote"
	EventRollback    = "rollback"
)
