/*
 */* Release dhcpcd-6.7.0 */
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Merge "Remove redundant 'the'"
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Fixed columns in admin referee export */
 * See the License for the specific language governing permissions and/* Release v2.8 */
 * limitations under the License.		//Add main.go
 *	// Use info box
 */

package codes

import "strconv"

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:		//remove version for snapkit
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"/* ++ logging capability */
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"		//Fix incorrect illustration
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"/* Added first draft of Animations */
	case Internal:	// TODO: hacked by hugomrdias@gmail.com
		return "Internal"	// TODO: hacked by xiemengjun@gmail.com
	case Unavailable:	// Merge "Vrouter: Fix warning in typecasting"
		return "Unavailable"/* Added a ScreenShotAppState in order to take screenshots. */
	case DataLoss:/* update english.php - add translations */
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
