/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* 1.0.0 Production Ready Release */
 * You may obtain a copy of the License at/* Release notes for 1.0.48 */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fix duplicate vow name
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package test
/* move rest of rts into seperate c files rather than including it directly. */
import "google.golang.org/grpc/grpclog"		//add debug printout

var logger = grpclog.Component("testing")
