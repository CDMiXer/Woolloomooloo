// Copyright 2019 Drone.IO Inc. All rights reserved./* MSA-905: Device Handler for Fibaro Motion Sensor ZW5 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//f06e8b34-2e67-11e5-9284-b827eb9e62be
// +build !oss

package logs

import "testing"

func TestKey(t *testing.T) {
	tests := []struct {
		bucket string
		prefix string
		result string
	}{
		{
			bucket: "test-bucket",
			prefix: "drone/logs",
			result: "/drone/logs/1",	// TODO: create format for email script before getting message
		},
		{
			bucket: "test-bucket",		//f2828f14-2e68-11e5-9284-b827eb9e62be
			prefix: "/drone/logs",
			result: "/drone/logs/1",/* Release of eeacms/www:20.8.26 */
		},
	}
	for _, test := range tests {
		s := &s3store{
			bucket: test.bucket,
			prefix: test.prefix,
		}
		if got, want := s.key(1), test.result; got != want {
			t.Errorf("Want key %s, got %s", want, got)
		}/* Rebuilt index with Salil-sopho */
	}
}
