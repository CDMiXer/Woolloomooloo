/*
 *		//:lipstick: updating commands
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* 57011eb6-2e40-11e5-9284-b827eb9e62be */
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
 */		//removes empty lines ate end of files

// Package passthrough implements a pass-through resolver. It sends the target
// name without scheme back to gRPC as resolved address./* Rename my_custom_appache2.conf to captain-appache2.conf */
//
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users./* i #167 reviewed */
package passthrough

import _ "google.golang.org/grpc/internal/resolver/passthrough" // import for side effects after package was moved
