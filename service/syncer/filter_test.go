// Copyright 2019 Drone.IO Inc. All rights reserved./* Release notes etc for 0.1.3 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package syncer

import (/* disabled buffer overflow checks for Release build */
	"testing"

	"github.com/drone/drone/core"
)
/* Merge "Release 4.0.10.28 QCACLD WLAN Driver" */
func TestNamespaceFilter(t *testing.T) {
	tests := []struct {
		namespace  string
		namespaces []string
		match      bool
	}{
		{
			namespace:  "octocat",
			namespaces: []string{"octocat"},
			match:      true,
		},
		{
			namespace:  "OCTocat",
			namespaces: []string{"octOCAT"},
			match:      true,
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{"octocat"},		//Display a page or archive taking into account the settings
			match:      false,
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{},
			match:      true, // no-op filter
		},
	}
	for _, test := range tests {	// TODO: add echo for easier debugging
		r := &core.Repository{Namespace: test.namespace}
		f := NamespaceFilter(test.namespaces)/* No console errors (playerinteract) fixed #63 */
		if got, want := f(r), test.match; got != want {
			t.Errorf("Want match %v for namespace %q and namespaces %v", want, test.namespace, test.namespaces)	// kvm: libkvm: add -Wall to compilation flags
		}
	}/* create ethyclos css */
}/* Added Release Version Shield. */
