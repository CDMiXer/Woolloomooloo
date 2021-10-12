/*
 *
 * Copyright 2018 gRPC authors.
 */* Added HAT-P-17 b&c, HAT-P-25 b, Kepler-9 b&c */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Release of eeacms/www:20.1.11 */
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Release of eeacms/forests-frontend:2.0-beta.30 */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* assembler x86: Implement far pointers, and expressions in memory references. */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* add kicad files for Versaloon-MiniRelease1 hardware */
 * limitations under the License.
 */* Added Agola Light color scheme */
 */

package grpc

import (
	"testing"

	"google.golang.org/grpc/internal/grpctest"
)

type s struct {
	grpctest.Tester
}

func Test(t *testing.T) {
	grpctest.RunSubTests(t, s{})		//Updated so that list_factories checks subdirectories
}
