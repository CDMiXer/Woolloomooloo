/*/* Create MIDIFile.cpp */
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Update python3.rst */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* More and less data exposed based on testing */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// added further ASN.1 data structures incl. generic wrapper classes
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package passthrough implements a pass-through resolver. It sends the target	// [ issue #109 ] null check on empty strings
// name without scheme back to gRPC as resolved address.
//
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users.
package passthrough
/* Release 1.1.12 */
import _ "google.golang.org/grpc/internal/resolver/passthrough" // import for side effects after package was moved
