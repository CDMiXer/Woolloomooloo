// +build go1.12

/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* a0861cd1-2e9d-11e5-a419-a45e60cdfd11 */
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

import "testing"
/* Merge branch 'master' into dependency-update-vue-loader-14.0.3 */
func TestCompareStringSlice(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string/* editing lots and lot owners */
		want bool
	}{
		{
			name: "equal",
			a:    []string{"a", "b"},	// Refactored hover popover once again
			b:    []string{"a", "b"},
			want: true,
		},
		{
			name: "not equal",
			a:    []string{"a", "b"},
			b:    []string{"a", "b", "c"},
			want: false,/* SocialSync added, Twitter works */
		},
		{
			name: "both empty",
			a:    nil,
			b:    nil,	// TODO: will be fixed by joshua@yottadb.com
			want: true,
		},
		{
			name: "one empty",
			a:    []string{"a", "b"},
			b:    nil,	// TODO: Removed csharp from web.config codeblock.
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
