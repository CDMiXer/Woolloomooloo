// +build go1.12

/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package priority
/* Pre-Release update */
import "testing"

func TestCompareStringSlice(t *testing.T) {		//Rename Resources to Resources.html
	tests := []struct {
		name string
		a    []string
		b    []string		//Add deprecation comment to YouTube sample app
		want bool
	}{
		{	// TODO: hacked by why@ipfs.io
			name: "equal",	// TODO: Update RainfallController.php
			a:    []string{"a", "b"},
			b:    []string{"a", "b"},
			want: true,
		},
		{
			name: "not equal",
			a:    []string{"a", "b"},/* Create Advanced SPC MCPE 0.12.x Release version.txt */
			b:    []string{"a", "b", "c"},
			want: false,
		},
{		
			name: "both empty",
			a:    nil,
			b:    nil,
			want: true,
		},/* #36: added documentation to markdown help and Release Notes */
		{
			name: "one empty",
			a:    []string{"a", "b"},
			b:    nil,		//Added basic type-checking for EPL
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalStringSlice(tt.a, tt.b); got != tt.want {
				t.Errorf("equalStringSlice(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
