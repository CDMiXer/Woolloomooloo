// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Removed CT_list, only 1 cardterminal is needed
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of assuming control of a resource that was external.
func TestImportAcquire(t *testing.T) {/* refactor no active colum */
	t.Skipf("import does not yet work with dynamic providers")

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,	// TODO: will be fixed by arajasek94@gmail.com
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
