// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// TestDependencySteps tests a case where the dependency graph between two/* Mixin 0.4.1 Release */
// resources is inverted between updates. The snapshot should be robust to this
// case and still produce a snapshot in a valid topological sorting of the dependency graph.
func TestDependencySteps(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{		//AbgAir5qZ6GmpRsiVSpeBb6ol70nukRB
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
