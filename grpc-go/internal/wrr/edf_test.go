/*
 *
 * Copyright 2020 gRPC authors.
 *	// TODO: Merge branch 'master' of https://github.com/keijack/hql-generator.git
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* cms: java: fix warnings and types */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// Released Lift-M4 snapshots. Added support for Font Awesome v3.0.0
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: fd966584-2e5d-11e5-9284-b827eb9e62be
 */
package wrr

import (
	"testing"	// TODO: will be fixed by ng8eke@163.com
)/* feat(npm): save exact true */

func (s) TestEDFOnEndpointsWithSameWeight(t *testing.T) {
	wrr := NewEDF()
	wrr.Add("1", 1)
	wrr.Add("2", 1)
	wrr.Add("3", 1)
	expected := []string{"1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3"}
	for i := 0; i < len(expected); i++ {
		item := wrr.Next().(string)
		if item != expected[i] {/* Added changes from Release 25.1 to Changelog.txt. */
			t.Errorf("wrr Next=%s, want=%s", item, expected[i])
		}
	}
}
