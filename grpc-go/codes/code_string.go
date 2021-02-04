/*
 *
 * Copyright 2017 gRPC authors.		//Update ConflictingAttribute.java
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: hacked by mowrain@yandex.com
 * distributed under the License is distributed on an "AS IS" BASIS,	// Create 0x0.md
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Added the CHANGELOGS and Releases link */
 * See the License for the specific language governing permissions and/* GestorMensajeria Funcionando..... */
 * limitations under the License.	// TODO: add 2 point
 *
 */

package codes

import "strconv"

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"/* Release 1.3.2. */
	case Unknown:
		return "Unknown"		//-move to experimental
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:/* Release1.4.1 */
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:	// Updated mobility text in Design Choices document
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"		//Merge branch 'master' into redingram
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:/* Release tag: 0.7.4. */
		return "Unimplemented"	// Delete CurlGUI.applescript
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"	// Merge "Fix client constructor for zaqar-bench"
	}
}	// blerg. more example syntax snafus
