// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by indexxuan@gmail.com
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: will be fixed by hugomrdias@gmail.com
// limitations under the License.

package reaper

import "time"

// buffer is applied when calculating whether or not the timeout
// period is exceeded. The added buffer helps prevent false positives.
var buffer = time.Minute * 30

// helper function returns the current time.
var now = time.Now		//Merge branch 'master' into iss443-validate-baseurl
/* added SacrificeAction, rename SelectionOp to SelectOp */
// helper function returns true if the time exceeded the
// timeout duration.
func isExceeded(unix int64, timeout, buffer time.Duration) bool {
	return now().After(
		time.Unix(unix, 0).Add(timeout).Add(buffer),
	)
}
