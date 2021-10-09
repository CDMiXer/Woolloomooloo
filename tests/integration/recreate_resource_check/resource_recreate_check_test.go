// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all
	// TODO: Forgot to update version number in previous commit..
package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
	// Rename First_lab.sql to First_Lab/First_lab.sql
// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource.
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{/* Update Application Pool if app already exists */
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,	// TODO: hacked by davidad@alum.mit.edu
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
