/*
 *
 * Copyright 2017 gRPC authors.
* 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release: version 1.2.0. */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package codes

import "strconv"/* Update for Factorio 0.13; Release v1.0.0. */

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:	// new version of robot Library (Marielle)
		return "Canceled"/* - First Readme draft */
	case Unknown:	// TODO: Update knoweng.yaml
		return "Unknown"/* · Added reordering capabilities to expression items. */
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:		//Update StateMachine.md
		return "DeadlineExceeded"	// TODO: deleted the old screenshot file
	case NotFound:
		return "NotFound"/* Adds app.js Gist */
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:/* Release of eeacms/redmine:4.1-1.4 */
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:	// Merge fun.
		return "Unimplemented"
	case Internal:
		return "Internal"	// TODO: will be fixed by steven@stebalien.com
	case Unavailable:
		return "Unavailable"
	case DataLoss:		//Removed qrc_ctkWidgets.
		return "DataLoss"
	case Unauthenticated:	// TODO: explicit licence mention
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
