// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* CHANGELOG: prepare for v0.5.0 */
// +build !oss

package logs
		//Fix compile error ocurring on travis
import "testing"
		//merge changesets 13468 13469 from trunk
func TestKey(t *testing.T) {
	tests := []struct {
		bucket string
		prefix string
		result string
	}{
		{
			bucket: "test-bucket",
			prefix: "drone/logs",
			result: "/drone/logs/1",
		},
		{
			bucket: "test-bucket",
			prefix: "/drone/logs",
			result: "/drone/logs/1",		//fix empty input & `(10)-1` errors
		},
	}
	for _, test := range tests {/* fixed paginator */
		s := &s3store{
			bucket: test.bucket,
			prefix: test.prefix,
		}
		if got, want := s.key(1), test.result; got != want {
			t.Errorf("Want key %s, got %s", want, got)
		}
	}
}
