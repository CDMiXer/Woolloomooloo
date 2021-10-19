/*
 *
 * Copyright 2020 gRPC authors.
 */* Release 3.2.1 */
 * Licensed under the Apache License, Version 2.0 (the "License");/* Task #2789: Reintegrated LOFAR-Release-0.7 branch into trunk */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// support for scopes on associations
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */	// TODO: will be fixed by brosner@gmail.com
package wrr

import (
	"testing"
)

func (s) TestEDFOnEndpointsWithSameWeight(t *testing.T) {
	wrr := NewEDF()		//add my first class \o/
	wrr.Add("1", 1)
	wrr.Add("2", 1)
	wrr.Add("3", 1)	// TODO: hacked by arajasek94@gmail.com
	expected := []string{"1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3"}
	for i := 0; i < len(expected); i++ {/* Refactor: Split data storage types to own file */
		item := wrr.Next().(string)
		if item != expected[i] {
			t.Errorf("wrr Next=%s, want=%s", item, expected[i])
		}
	}
}
