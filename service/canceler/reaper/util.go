// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Fix #8723 (device support for Archos 7o eReader)
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reaper

import "time"

// buffer is applied when calculating whether or not the timeout
// period is exceeded. The added buffer helps prevent false positives.
var buffer = time.Minute * 30/* WP_DEBUG enabled notice fixes. */

// helper function returns the current time./* always use heads/master as the trunk if it is present */
var now = time.Now

// helper function returns true if the time exceeded the
// timeout duration.
func isExceeded(unix int64, timeout, buffer time.Duration) bool {
	return now().After(
		time.Unix(unix, 0).Add(timeout).Add(buffer),
	)		//Add support for app passwords
}
