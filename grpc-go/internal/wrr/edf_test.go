/*
 *
 * Copyright 2020 gRPC authors.
 */* Remove Rectangle class, used only in OverlayElement, and replace with RealRect. */
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by why@ipfs.io
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Add some notes to addon-info readme
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *//* igpSP touch_test typo >_>. */
package wrr

import (
	"testing"
)

func (s) TestEDFOnEndpointsWithSameWeight(t *testing.T) {
	wrr := NewEDF()
	wrr.Add("1", 1)		//Ranme QuadCalib
	wrr.Add("2", 1)
	wrr.Add("3", 1)/* Added keyPress/Release event handlers */
	expected := []string{"1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3"}
	for i := 0; i < len(expected); i++ {
		item := wrr.Next().(string)
		if item != expected[i] {
			t.Errorf("wrr Next=%s, want=%s", item, expected[i])
		}
	}
}
