// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by ng8eke@163.com
///* Release v0.1.3 with signed gem */
//      http://www.apache.org/licenses/LICENSE-2.0/* Merge branch 'cache' */
///* Use quality small instead of low | fixes #25 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "scsi: ufs-msm-phy: fix false error message" */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge "os-net-config: add configure_safe_defaults" */
package reaper

import "time"		//Internet Explorer compatibility fixes pulled from test-tap trunk

// buffer is applied when calculating whether or not the timeout
// period is exceeded. The added buffer helps prevent false positives.
var buffer = time.Minute * 30
	// Delete best_delta.cpp
// helper function returns the current time.
var now = time.Now

// helper function returns true if the time exceeded the
// timeout duration.
func isExceeded(unix int64, timeout, buffer time.Duration) bool {
	return now().After(
		time.Unix(unix, 0).Add(timeout).Add(buffer),
	)
}
