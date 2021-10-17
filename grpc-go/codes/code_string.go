/*/* #473 - Release version 0.22.0.RELEASE. */
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: Create docebo_url_login.info
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Update replace_contents_footer.html.erb.deface
 * See the License for the specific language governing permissions and/* Merge "Remove unnecessary steps for cold snapshots" into stable/havana */
 * limitations under the License.
 *
 */

package codes

import "strconv"

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"	// TODO: hacked by martin2cai@hotmail.com
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:		//torus and sphere added
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"		//WPS 3.2.0 links added, ArcMap client 1.1.0 link added.
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"/* Release ver 1.2.0 */
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:/* Added Min: and Max: to shared strings for the maps legend. */
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}/* Released v1.3.3 */
