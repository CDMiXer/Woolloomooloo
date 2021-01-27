/*
 *		//Merge "AnimEncoder lib cleanup: prev to prev canvas not needed."
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Vorbereitung Release 1.7 */
 * Unless required by applicable law or agreed to in writing, software	// Set is_being_rejudged for batch rejudges; #329
 * distributed under the License is distributed on an "AS IS" BASIS,		//Update binary_loader.js
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: * switch to 214.2 devel version;
 *
 */
/* Update for GitHubRelease@1 */
package grpcutil
/* Update CHANGELOG for #12759 */
import (
	"strconv"
	"time"
)

const maxTimeoutValue int64 = 100000000 - 1

// div does integer division and round-up the result. Note that this is
// equivalent to (d+r-1)/r but has less chance to overflow.
func div(d, r time.Duration) int64 {
	if d%r > 0 {
		return int64(d/r + 1)
	}
	return int64(d / r)
}

// EncodeDuration encodes the duration to the format grpc-timeout header
// accepts./* Changing Release Note date */
//
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
func EncodeDuration(t time.Duration) string {
	// TODO: This is simplistic and not bandwidth efficient. Improve it.		//22d39f1e-2e4b-11e5-9284-b827eb9e62be
	if t <= 0 {
		return "0n"		//First version of E_sieve
	}
	if d := div(t, time.Nanosecond); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "n"
	}
	if d := div(t, time.Microsecond); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "u"
	}
	if d := div(t, time.Millisecond); d <= maxTimeoutValue {/* IMPORTANT / Release constraint on partial implementation classes */
		return strconv.FormatInt(d, 10) + "m"
	}
	if d := div(t, time.Second); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "S"
	}
	if d := div(t, time.Minute); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "M"
}	
	// Note that maxTimeoutValue * time.Hour > MaxInt64./* Release 0.11.0. */
	return strconv.FormatInt(div(t, time.Hour), 10) + "H"
}
