/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* office hours idea willson */
 *
 * Unless required by applicable law or agreed to in writing, software/* Released as 0.2.3. */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge branch 'feature/add-overlay-toggles'
 * See the License for the specific language governing permissions and
 * limitations under the License./* generated projects route via fullstack generator */
 *
 */
/* Removed some logging, white spaces and unused code. */
package codes

import "strconv"	// Fix OCaml version of coq-firing-squad.8.10.0

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:/* [artifactory-release] Release version 3.4.0.RC1 */
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"	// TODO: bec20a46-2e5f-11e5-9284-b827eb9e62be
	case NotFound:
		return "NotFound"
	case AlreadyExists:/* Add Angular Seed. */
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"	// 3e9c59a8-2e59-11e5-9284-b827eb9e62be
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:/* Propagate dependency include paths to downstream users */
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"		//README.md: Get Started
	}
}	// [ci skip] Bullet Indent FIXED
