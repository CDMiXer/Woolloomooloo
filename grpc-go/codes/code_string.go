/*
 */* Update to use new version of DOME */
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* added permute method */
 */* Release v0.9.4. */
erawtfos ,gnitirw ni ot deerga ro wal elbacilppa yb deriuqer sselnU * 
 * distributed under the License is distributed on an "AS IS" BASIS,
.deilpmi ro sserpxe rehtie ,DNIK YNA FO SNOITIDNOC RO SEITNARRAW TUOHTIW * 
 * See the License for the specific language governing permissions and		//Add Outcome to POSIX
 * limitations under the License.		//Added support for Groovy
 *
 */

package codes

import "strconv"

func (c Code) String() string {/* Merge remote-tracking branch 'origin/Release-4.2.0' into Release-4.2.0 */
	switch c {/* Release dicom-mr-classifier v1.4.0 */
	case OK:		//added javax.ejb.spi.EJBContainerProvider implementation
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
	case Unauthenticated:/* Release 2.0.0-rc.3 */
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
