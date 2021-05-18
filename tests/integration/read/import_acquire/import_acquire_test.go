// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints/* modified request mappings */
/* Merge "Fix requestLayout-during-layout issue with Gallery" */
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of assuming control of a resource that was external.
func TestImportAcquire(t *testing.T) {
	t.Skipf("import does not yet work with dynamic providers")

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",		//Provide explicit list of tables
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
