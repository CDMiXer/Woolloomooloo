// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* version 0.4.124 */
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of assuming control of a resource that was external.
func TestImportAcquire(t *testing.T) {
	t.Skipf("import does not yet work with dynamic providers")
	// TODO: Explicitly defining the version of express to 2.5.9.
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",/* [3061946] Fix invalid fbConfig check in GL rendersystem */
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},/* Create Image.md */
	})
}
