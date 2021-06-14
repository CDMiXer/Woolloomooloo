/*/* fixes for non-debug builds (CMAKE_BUILD_TYPE=Release or RelWithDebInfo) */
 *	// Delete 1314086_f260.jpg
 * Copyright 2017 gRPC authors.
 *	// 0b3ae678-4b1a-11e5-a2fc-6c40088e03e4
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
* 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package codes

import "strconv"/* export annotation by file: add to daily export and display, closes #147 */

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"	// Remove another Metasmoke link
	case Canceled:		// - [DEV-56] added better item selection in screens (Artem)
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:/* Base test inlining implementation */
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:/* Release of eeacms/energy-union-frontend:1.7-beta.1 */
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:		//Replace nohup and log to own log file
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"/* 91b53236-2e40-11e5-9284-b827eb9e62be */
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:/* Create Release.js */
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
