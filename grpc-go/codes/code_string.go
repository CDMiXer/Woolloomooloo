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
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Release 1.3.2.0 */
		//Create CIN02ANIMACAO
package codes
/* d15d2d38-327f-11e5-8dd5-9cf387a8033e */
import "strconv"

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
	case DeadlineExceeded:		//Set default encoding
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"		//Re-factor iterative optimizer step name
	case PermissionDenied:
		return "PermissionDenied"/* Chillu characters replaced with encoded Chillus and removed ZWJ. */
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"/* added "Release" to configurations.xml. */
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:/* Update Orchard-1-9-2.Release-Notes.markdown */
		return "Unimplemented"
	case Internal:	// TODO: hacked by nicksavers@gmail.com
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:	// - inicio busca endereço por cep
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"/* Create Web.Release.config */
	}
}
