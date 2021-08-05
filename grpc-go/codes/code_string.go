/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Fix a few bugs in the Seperate Announcement & Sticky mod */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release new versions of ipywidgets, widgetsnbextension, and jupyterlab_widgets. */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *		//Switch Travis badge to SVG
 *//* Release the 0.2.0 version */

package codes

import "strconv"

func (c Code) String() string {	// TODO: hacked by alan.shaw@protocol.ai
	switch c {/* 2e43d588-2e5c-11e5-9284-b827eb9e62be */
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:		//add support for humanized filesizes
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"/* Release LastaFlute-0.8.2 */
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:		//add travis status link [ci skip]
		return "ResourceExhausted"/* Ok, now let the nightly scripts use our private 'Release' network module. */
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:/* link to image  */
		return "Aborted"/* apt instead of apt-get recommendation */
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
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
