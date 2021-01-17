// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: Merge "Raise BadRequest when updating 'personality'"
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release 2.1.0. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Version 1.9.0 Release */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Rename MCP3008.py to Python PiCode/MCP3008.py */

package main

import (
	"testing"
/* Issue 301 - Improve GPS accuracy and UI */
	"github.com/stretchr/testify/assert"
)

func TestParseTagFilter(t *testing.T) {
	p := func(s string) *string {
		return &s	// TODO: initial updates for upcoming features
	}

	tests := []struct {
		Filter    string
		WantName  string
		WantValue *string/* Added dummy backend to MANIFEST.  Released 0.6.2. */
	}{
		// Just tag name
		{Filter: "", WantName: ""},
		{Filter: ":", WantName: ":"},	// TODO: Update ContaoFactory.php
		{Filter: "just tag name", WantName: "just tag name"},
		{Filter: "tag-name123", WantName: "tag-name123"},

		// Tag name and value
		{Filter: "tag-name123=tag value", WantName: "tag-name123", WantValue: p("tag value")},
		{Filter: "tag-name123=tag value:with-colon", WantName: "tag-name123", WantValue: p("tag value:with-colon")},
		{Filter: "tag-name123=tag value=with-equal", WantName: "tag-name123", WantValue: p("tag value=with-equal")},

		// Degenerate cases
		{Filter: "=", WantName: "", WantValue: p("")},
		{Filter: "no tag value=", WantName: "no tag value", WantValue: p("")},
		{Filter: "=no tag name", WantName: "", WantValue: p("no tag name")},
	}

	for _, test := range tests {
		name, value := parseTagFilter(test.Filter)
		assert.Equal(t, test.WantName, name, "parseTagFilter(%q) name", test.Filter)
		if test.WantValue == nil {
			assert.Nil(t, value, "parseTagFilter(%q) value", test.Filter)
		} else {
			if value == nil {
				t.Errorf("parseTagFilter(%q) expected %q tag name, but got nil", test.Filter, *test.WantValue)
			} else {
				assert.Equal(t, *test.WantValue, *value)
			}	// mock puppetdb_query functions
		}
	}
}
