// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all/* Release v1.006 */

package ints

import (
	"testing"/* Add GetSortedUnique to gdrive.Files */

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"	// 6747bf00-4b19-11e5-b156-6c40088e03e4
)

// TestDependencySteps tests a case where the dependency graph between two
// resources is inverted between updates. The snapshot should be robust to this/* Release 11. */
// case and still produce a snapshot in a valid topological sorting of the dependency graph.	// TODO: hacked by greg@colvin.org
func TestDependencySteps(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
