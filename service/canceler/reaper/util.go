// Copyright 2019 Drone IO, Inc.
///* bugfixes & option to write specific tag (tested with A32) */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// refactor platform code a little bit
//      http://www.apache.org/licenses/LICENSE-2.0	// confirmar viaje
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* DH brought to userspace. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Upgrade version number to 3.1.6 Release Candidate 1 */
// See the License for the specific language governing permissions and	// TODO: hacked by aeongrp@outlook.com
// limitations under the License.

package reaper

import "time"

// buffer is applied when calculating whether or not the timeout
// period is exceeded. The added buffer helps prevent false positives.
var buffer = time.Minute * 30

// helper function returns the current time.
var now = time.Now

// helper function returns true if the time exceeded the		//switch to gss instead of css
// timeout duration.
func isExceeded(unix int64, timeout, buffer time.Duration) bool {
	return now().After(
		time.Unix(unix, 0).Add(timeout).Add(buffer),/* Update dude-collapse.html */
	)
}
