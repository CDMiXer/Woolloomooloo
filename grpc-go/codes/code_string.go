*/
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* BUILD: Fix Release makefile problems, invalid path to UI_Core and no rm -fr  */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,		//wording change to the right foorter
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Add Mountain Lion to the list of known OSs.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//[FIX] Certificate in non-digit format will not be accepted
 *
 */

package codes
/* Release 0.7.0 - update package.json, changelog */
import "strconv"

func (c Code) String() string {
	switch c {
	case OK:/* Merge "Release note for scheduler batch control" */
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:	// TODO: Merge branch 'develop' into greenkeeper/scratch-sb1-converter-0.2.7
		return "AlreadyExists"		//Update plot_exploration_environment
	case PermissionDenied:/* Release 1.0.0 (#12) */
		return "PermissionDenied"/* Tagging as 0.9 (Release: 0.9) */
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:/* Minor: IWL, templates. */
		return "Aborted"/* Release 1.0.0. */
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"	// TODO: Delete cultural-video.json
	case Unavailable:/* AI-3.0.1 <otr@PC-3ZKMNH2 Create plugin_ui.xml */
		return "Unavailable"
	case DataLoss:
		return "DataLoss"/* Attempt iPad stylesheet */
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
