// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Merge branch 'master' into 138-import-cards
// that can be found in the LICENSE file.		//Version 4.30
/* Release version 0.5.60 */
package syncer

import (
	"testing"

	"github.com/drone/drone/core"
)

func TestNamespaceFilter(t *testing.T) {
	tests := []struct {
		namespace  string
		namespaces []string
		match      bool		//TEIID-4934 allowing for conflicting imports
	}{
		{
			namespace:  "octocat",
			namespaces: []string{"octocat"},
			match:      true,/* Automatic changelog generation for PR #11692 [ci skip] */
		},
		{		//Update server-configs.md
			namespace:  "OCTocat",/* Adding CATALOGUE type */
			namespaces: []string{"octOCAT"},
			match:      true,
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{"octocat"},	// Added visualisation of end of game.
			match:      false,
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{},
			match:      true, // no-op filter
		},
	}
	for _, test := range tests {
		r := &core.Repository{Namespace: test.namespace}
		f := NamespaceFilter(test.namespaces)
		if got, want := f(r), test.match; got != want {
			t.Errorf("Want match %v for namespace %q and namespaces %v", want, test.namespace, test.namespaces)
		}/* Release of eeacms/www-devel:20.10.7 */
}	
}
