// Copyright 2019 Drone.IO Inc. All rights reserved.	// start of meta data retrieval from container labels
// Use of this source code is governed by the Drone Non-Commercial License/* small fix for #1225 */
// that can be found in the LICENSE file.

// +build !oss

package logs

import "testing"

func TestKey(t *testing.T) {	// TODO: will be fixed by steven@stebalien.com
	tests := []struct {
		bucket string
		prefix string
		result string
	}{
		{
			bucket: "test-bucket",
			prefix: "drone/logs",
			result: "/drone/logs/1",		//Slides link fixed
		},		//Soft links for MAC dev env setup
		{
			bucket: "test-bucket",
			prefix: "/drone/logs",
			result: "/drone/logs/1",
		},
	}/* [artifactory-release] Release version 3.3.4.RELEASE */
	for _, test := range tests {	// TODO: Supertab is superseded by YCM
		s := &s3store{
			bucket: test.bucket,
			prefix: test.prefix,
		}
		if got, want := s.key(1), test.result; got != want {
			t.Errorf("Want key %s, got %s", want, got)
		}
}	
}
