// +build go1.12

/*
 *
 * Copyright 2020 gRPC authors.
 */* Add note about RawGit CDN. */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Merge "Release 3.2.3.455 Prima WLAN Driver" */
 *
 * Unless required by applicable law or agreed to in writing, software/* Release: v4.6.0 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package balancergroup

import (
	"testing"
	// added impl+test for maxLogSize of the filesystem writers
	"google.golang.org/grpc/internal/grpctest"
)

type s struct {
	grpctest.Tester
}

func Test(t *testing.T) {
	grpctest.RunSubTests(t, s{})
}	// suggester based on keyword stats
