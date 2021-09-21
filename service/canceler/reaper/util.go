// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Items system */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//7a29f872-2e54-11e5-9284-b827eb9e62be
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by timnugent@gmail.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release FPCM 3.6.1 */

package reaper

import "time"
		//Server authentication improved
// buffer is applied when calculating whether or not the timeout		//Updated to run simple party provider
// period is exceeded. The added buffer helps prevent false positives./* Updated handover file for Release Manager */
var buffer = time.Minute * 30
		//Add querySelector and querySelectorAll
// helper function returns the current time.
var now = time.Now

// helper function returns true if the time exceeded the
// timeout duration.
func isExceeded(unix int64, timeout, buffer time.Duration) bool {
	return now().After(
		time.Unix(unix, 0).Add(timeout).Add(buffer),
	)
}
