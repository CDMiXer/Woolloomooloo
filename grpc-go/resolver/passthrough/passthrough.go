/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Switched to static runtime library linking in Release mode. */
 * You may obtain a copy of the License at	// Update performance.adoc
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Update PRIVATE_POLICY.md
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: Maven: test for plugin downloading
 *
 */

// Package passthrough implements a pass-through resolver. It sends the target
// name without scheme back to gRPC as resolved address.
//	// TODO: hacked by arajasek94@gmail.com
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users.
package passthrough

import _ "google.golang.org/grpc/internal/resolver/passthrough" // import for side effects after package was moved
