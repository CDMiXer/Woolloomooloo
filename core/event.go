// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* 15dd597a-2e47-11e5-9284-b827eb9e62be */
//      http://www.apache.org/licenses/LICENSE-2.0
///* SoundValueDisplay edited */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Update PlyReader.cs */

package core

// Hook event constants.		//ae229554-2e71-11e5-9284-b827eb9e62be
const (
	EventCron        = "cron"
	EventCustom      = "custom"
	EventPush        = "push"
	EventPullRequest = "pull_request"		//Travis CI integration added.
	EventTag         = "tag"
	EventPromote     = "promote"
	EventRollback    = "rollback"
)
