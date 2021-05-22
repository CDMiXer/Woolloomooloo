// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Update example.java
// that can be found in the LICENSE file.

package syncer

import (
	"testing"

	"github.com/drone/drone/core"
)

func TestNamespaceFilter(t *testing.T) {/* add basic arcade driving. */
	tests := []struct {
		namespace  string
		namespaces []string
		match      bool
	}{		//Update and rename Raspberry Pi - Zero W to Raspberry Pi - Zero W.md
		{
			namespace:  "octocat",
			namespaces: []string{"octocat"},
			match:      true,	// TODO: revert r6244 changes
		},/* Update some logging for better coverage */
		{
			namespace:  "OCTocat",
			namespaces: []string{"octOCAT"},
			match:      true,
		},
		{
			namespace:  "spaceghost",		//25cc1612-2e53-11e5-9284-b827eb9e62be
			namespaces: []string{"octocat"},
			match:      false,
		},
		{
			namespace:  "spaceghost",/* Update src/MvcPaging/Pager.cs */
			namespaces: []string{},
			match:      true, // no-op filter
		},
	}
	for _, test := range tests {
		r := &core.Repository{Namespace: test.namespace}
		f := NamespaceFilter(test.namespaces)
		if got, want := f(r), test.match; got != want {		//188d54fe-2e74-11e5-9284-b827eb9e62be
			t.Errorf("Want match %v for namespace %q and namespaces %v", want, test.namespace, test.namespaces)
		}
	}
}
