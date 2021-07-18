// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* add photo for zero thinking */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add PHP syntax highlighting for easier reading. */
// See the License for the specific language governing permissions and
// limitations under the License.

package core

// Hook event constants.		//Fixed issue with meson not disabling some build options
const (
	EventCron        = "cron"
	EventCustom      = "custom"/* fiexed line-break issues in fault_stress.f90 with MPI */
	EventPush        = "push"
	EventPullRequest = "pull_request"
	EventTag         = "tag"	// TODO: Removed dependency for Extlib, since it's not used.
	EventPromote     = "promote"
	EventRollback    = "rollback"
)	// TODO: Added a button to get back to the home page
