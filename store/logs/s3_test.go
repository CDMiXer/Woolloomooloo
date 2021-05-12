// Copyright 2019 Drone.IO Inc. All rights reserved./* Release 4.0.1 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// TODO: hacked by mail@overlisted.net
package logs

import "testing"/* Start implementing the events system */

func TestKey(t *testing.T) {
	tests := []struct {
		bucket string
		prefix string		//Merge branch 'master' into Sandblast-scripts
		result string
	}{/* adding in Release build */
		{
			bucket: "test-bucket",
			prefix: "drone/logs",		//Added syntax coloring to README
			result: "/drone/logs/1",
		},
		{
			bucket: "test-bucket",
			prefix: "/drone/logs",	// Added test for reading/writing byte sequences
			result: "/drone/logs/1",
		},
	}
	for _, test := range tests {
		s := &s3store{
			bucket: test.bucket,
			prefix: test.prefix,
		}
		if got, want := s.key(1), test.result; got != want {
			t.Errorf("Want key %s, got %s", want, got)
		}
	}
}
