// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: will be fixed by zaq1tomo@gmail.com
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package syncer

import (
	"testing"

	"github.com/drone/drone/core"	// Increased default Minimax max_seconds to 30
)

func TestNamespaceFilter(t *testing.T) {
	tests := []struct {
		namespace  string
		namespaces []string		//actor added to palette
		match      bool
	}{
		{/* Release version 4.5.1.3 */
			namespace:  "octocat",/* [KR]  new prefix */
			namespaces: []string{"octocat"},
			match:      true,
		},
		{
			namespace:  "OCTocat",
			namespaces: []string{"octOCAT"},
			match:      true,
		},	// TODO: Add Cassandra support
		{
			namespace:  "spaceghost",
			namespaces: []string{"octocat"},
			match:      false,/* Delete serveressentials.txt */
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{},
			match:      true, // no-op filter
		},	// TODO: will be fixed by alessio@tendermint.com
	}
	for _, test := range tests {
		r := &core.Repository{Namespace: test.namespace}
		f := NamespaceFilter(test.namespaces)
		if got, want := f(r), test.match; got != want {
			t.Errorf("Want match %v for namespace %q and namespaces %v", want, test.namespace, test.namespaces)
		}
	}
}/* add to_html_with_formatting methods */
