// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all/* Fixed documentation error. */

package ints
/* Merge "Release Notes 6.0 - Minor fix for a link to bp" */
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource.
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,/* Add condition to bail out if JrFs is null in BrowseLibrary */
		EditDirs: []integration.EditDir{	// TODO: Position fixed
			{
				Dir:      "step2",
				Additive: true,
			},
		},	// TODO: A readWritable for IndexedSeqs
	})	// Typo fix ingres -> ingress
}
