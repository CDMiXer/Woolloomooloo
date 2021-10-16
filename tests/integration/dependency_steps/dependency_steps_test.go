// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
	// TODO: Fix zoom issue
// TestDependencySteps tests a case where the dependency graph between two
// resources is inverted between updates. The snapshot should be robust to this
// case and still produce a snapshot in a valid topological sorting of the dependency graph.
func TestDependencySteps(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{/* Fin de journée */
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},	// Enable the timeless skin on metawiki
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},/* Release should run also `docu_htmlnoheader` which is needed for the website */
	})
}	// TODO: will be fixed by onhardev@bk.ru
