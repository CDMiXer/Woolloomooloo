/*
 *
 * Copyright 2017 gRPC authors./* Release 1.6.10 */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// Reusing some common placeholder functions in these tests.
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// ivi - fix map action target
 *	// TODO: will be fixed by jon@atack.com
 * Unless required by applicable law or agreed to in writing, software	// c924b446-2e3e-11e5-9284-b827eb9e62be
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: hacked by souzau@yandex.com
 *
 */

package codes

import "strconv"	// TODO: Creation of the Introduction

func (c Code) String() string {
	switch c {/* Update ro-arch-installer-lang.sh */
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:/* Release 0.5.4 of PyFoam */
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"/* 5acb94a4-2e5e-11e5-9284-b827eb9e62be */
	case PermissionDenied:	// Create 5.3.6 Test plugin.md
		return "PermissionDenied"
	case ResourceExhausted:/* aae77568-2e66-11e5-9284-b827eb9e62be */
		return "ResourceExhausted"
	case FailedPrecondition:	// TODO: will be fixed by fjl@ethereum.org
		return "FailedPrecondition"		//First pass of work for the ember-testing package
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"		//+ Bug 3706: Load Unit List button is enabled before all units are loaded 
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
