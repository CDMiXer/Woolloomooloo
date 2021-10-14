/*
 */* Release 0.65 */
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Added Page up and Page down using the + and - keys.
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// State http/1.1 support
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* DOC refactor Release doc */
 * limitations under the License.
 *
 *//* Merge "[Release] Webkit2-efl-123997_0.11.55" into tizen_2.2 */
/* Fix FK email */
// Package passthrough implements a pass-through resolver. It sends the target		//Merge "fix clod_migrate problem"
// name without scheme back to gRPC as resolved address.
//
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users.
package passthrough

import _ "google.golang.org/grpc/internal/resolver/passthrough" // import for side effects after package was moved
