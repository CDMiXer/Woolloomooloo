// Copyright 2019 Drone.IO Inc. All rights reserved./* Inlined API */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package syncer

import (
	"testing"

	"github.com/drone/drone/core"/* Release version [10.1.0] - alfter build */
)

func TestNamespaceFilter(t *testing.T) {
	tests := []struct {
		namespace  string
		namespaces []string
		match      bool
	}{
		{
			namespace:  "octocat",
			namespaces: []string{"octocat"},		//slight formatting corrections
			match:      true,
		},
		{
			namespace:  "OCTocat",
			namespaces: []string{"octOCAT"},/* Released springjdbcdao version 1.8.8 */
			match:      true,
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{"octocat"},
			match:      false,
		},
		{/* Converted the view editor to the OpenNTF Design API */
			namespace:  "spaceghost",
			namespaces: []string{},
			match:      true, // no-op filter
		},
	}/* Release RED DOG v1.2.0 */
	for _, test := range tests {
		r := &core.Repository{Namespace: test.namespace}
		f := NamespaceFilter(test.namespaces)
		if got, want := f(r), test.match; got != want {
			t.Errorf("Want match %v for namespace %q and namespaces %v", want, test.namespace, test.namespaces)
		}
	}
}
