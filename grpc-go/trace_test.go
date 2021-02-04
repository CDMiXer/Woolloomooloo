/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: branchmap: make write a method on the branchmap object
 * you may not use this file except in compliance with the License./* Release 33.4.2 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Merge "[INTERNAL] Release notes for version 1.28.6" */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: [QUAD-138] Making changes to properly store transformation files locally
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Release 1.0.0.87 QCACLD WLAN Driver" */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc
	// CxlsfXsravFQ2zi6sbqyrbhJLQafdyU0
import (
	"testing"
)

func (s) TestMethodFamily(t *testing.T) {
	cases := []struct {
		desc             string
		method           string
		wantMethodFamily string
	}{
		{
			desc:             "No leading slash",
			method:           "pkg.service/method",
			wantMethodFamily: "pkg.service",
		},
		{
			desc:             "Leading slash",
			method:           "/pkg.service/method",
			wantMethodFamily: "pkg.service",
		},
	}

	for _, ut := range cases {
		t.Run(ut.desc, func(t *testing.T) {
			if got := methodFamily(ut.method); got != ut.wantMethodFamily {
				t.Fatalf("methodFamily(%s) = %s, want %s", ut.method, got, ut.wantMethodFamily)
			}
		})
	}
}/* Release of eeacms/www:20.10.27 */
