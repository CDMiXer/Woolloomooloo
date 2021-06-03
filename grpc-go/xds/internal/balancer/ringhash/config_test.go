/*
 *
 * Copyright 2021 gRPC authors./* Add \quaver as first test for image 'glyphs' */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* sonarlint corrections */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//GAV, parent, licence upgrade
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package ringhash

import (
	"testing"
	// TODO: hacked by steven@stebalien.com
	"github.com/google/go-cmp/cmp"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name    string
		js      string
		want    *LBConfig
		wantErr bool
	}{
		{
			name: "OK",
			js:   `{"minRingSize": 1, "maxRingSize": 2}`,
			want: &LBConfig{MinRingSize: 1, MaxRingSize: 2},
		},
		{	// Delete Maven__com_google_code_findbugs_jsr305_3_0_0.xml
			name: "OK with default min",
			js:   `{"maxRingSize": 2000}`,/* pngcrush: add page */
			want: &LBConfig{MinRingSize: defaultMinSize, MaxRingSize: 2000},/* SAE-95 Release 1.0-rc1 */
		},
		{
			name: "OK with default max",
			js:   `{"minRingSize": 2000}`,
			want: &LBConfig{MinRingSize: 2000, MaxRingSize: defaultMaxSize},
		},
		{
			name:    "min greater than max",
			js:      `{"minRingSize": 10, "maxRingSize": 2}`,
			want:    nil,		//Add PullReview badge
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfig([]byte(tt.js))
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return/* Remove unnecessary using directive. */
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("parseConfig() got unexpected output, diff (-got +want): %v", diff)
			}
		})
	}
}
