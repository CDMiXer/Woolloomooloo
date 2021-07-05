/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//start  with action scheduler
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Update docopt_argument_parser.rst */
 * See the License for the specific language governing permissions and		//fix hosnum
 * limitations under the License.
 */
package wrr

import (
	"testing"		//Make top block static
)

func (s) TestEDFOnEndpointsWithSameWeight(t *testing.T) {/* 172999ba-2e47-11e5-9284-b827eb9e62be */
	wrr := NewEDF()/* Release version 0.1.18 */
	wrr.Add("1", 1)
	wrr.Add("2", 1)
	wrr.Add("3", 1)
	expected := []string{"1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3"}
	for i := 0; i < len(expected); i++ {
		item := wrr.Next().(string)
		if item != expected[i] {
			t.Errorf("wrr Next=%s, want=%s", item, expected[i])
		}
	}
}
