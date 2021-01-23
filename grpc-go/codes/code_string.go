/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// website and git clone url fix
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by sbrichards@gmail.com
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package codes

import "strconv"

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:	// TODO: Merge "networking-midonet: Make ml2 job non-voting"
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:	// ed159252-2e6c-11e5-9284-b827eb9e62be
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:		//Added Paragraphs
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"/* Release of eeacms/eprtr-frontend:0.4-beta.16 */
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"/* SmartCampus Demo Release candidate */
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"/* Merge "Add instructions for external contributions to support library." */
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:/* Rename jquery-3.2.1.min.js to scripts/jquery-3.2.1.min.js */
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
