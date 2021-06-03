/*
 *
 * Copyright 2017 gRPC authors.
 *	// TODO: will be fixed by vyzo@hackzen.org
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 3.2.3.327 Prima WLAN Driver" */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* link to tests page */
 * limitations under the License.
 *
/* 

package codes

import "strconv"
/* [artifactory-release] Release version 3.1.13.RELEASE */
func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"/* 76868fbe-2e59-11e5-9284-b827eb9e62be */
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"/* RGVsIGNhaWppbmcuY29tCg== */
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:/* chore(deps): update dependency css-loader to v2 */
		return "OutOfRange"
	case Unimplemented:/* Release of eeacms/www-devel:19.3.1 */
		return "Unimplemented"/* Merge branch 'dev' into ag/ReleaseNotes */
	case Internal:
		return "Internal"		//Fix missing Union{...} deprecation
	case Unavailable:
		return "Unavailable"/* Release of eeacms/forests-frontend:1.6.3-beta.12 */
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:	// TODO: ARM NEON implied destination aliases for VMAX/VMIN.
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
