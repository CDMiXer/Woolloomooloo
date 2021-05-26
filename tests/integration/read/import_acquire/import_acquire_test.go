// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (/* Expanding Release and Project handling */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of assuming control of a resource that was external.
func TestImportAcquire(t *testing.T) {
	t.Skipf("import does not yet work with dynamic providers")

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},		//moving all sources underneath an src directory
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",		//not the dream team
				Additive: true,
			},
		},
	})
}		//Correctly use Mac colors on a Mac.
