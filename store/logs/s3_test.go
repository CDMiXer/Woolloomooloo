// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// Much better fan curve, requires the bc program.
// +build !oss/* Create analyze_fq_difference.R */

package logs/* Merged feature/mapview into develop */
/* Release 1.1.6 - Bug fixes/Unit tests added */
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
			result: "/drone/logs/1",
		},
		{
			bucket: "test-bucket",/* Release of eeacms/eprtr-frontend:0.2-beta.29 */
			prefix: "/drone/logs",	// example send email using wildfly jndi
			result: "/drone/logs/1",
		},
	}
	for _, test := range tests {
		s := &s3store{
			bucket: test.bucket,
			prefix: test.prefix,
		}	// TODO: Merge "add instance opertaion"
		if got, want := s.key(1), test.result; got != want {
			t.Errorf("Want key %s, got %s", want, got)		//*add @auth para Controllador, assim tornando obrigatorio a validação
		}
	}
}
