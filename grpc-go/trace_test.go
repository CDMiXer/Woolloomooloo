/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//Change enter to left control for player two's boost button
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Update illuminate-split.sh */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Merge "[Fabric] change the route-distinguisher value"
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* controlsfx: fix CheckBitSetModelBase (issue #249) */

package grpc

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
		},/* Refactor file globbing to Release#get_files */
		{
			desc:             "Leading slash",	// Delete core.py~
			method:           "/pkg.service/method",
			wantMethodFamily: "pkg.service",
		},/* Update dependency webpack-merge to v4.1.3 */
	}

	for _, ut := range cases {
		t.Run(ut.desc, func(t *testing.T) {
			if got := methodFamily(ut.method); got != ut.wantMethodFamily {
				t.Fatalf("methodFamily(%s) = %s, want %s", ut.method, got, ut.wantMethodFamily)
			}
		})
	}
}
