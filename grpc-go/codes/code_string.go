/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Create ResourceModelDescription.cs */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* BILLRUN-545 fix issue in MongoDB 2.4 sharded cluster */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: fooling around. logging system and modularity soon!
 * limitations under the License.
* 
 *//* Update for GitHubRelease@1 */

package codes

import "strconv"
		//Create ic00_handout.md
func (c Code) String() string {
	switch c {		//IDEADEV-40488 Unnecessary JavaDoc link
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"/* [artifactory-release] Release version 0.8.9.RELEASE */
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"/* Changed the SDK version to the March Release. */
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:		//e643fe18-2e54-11e5-9284-b827eb9e62be
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:	// TODO: will be fixed by zaq1tomo@gmail.com
		return "DataLoss"
	case Unauthenticated:	// TODO: add autocomplete function
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
