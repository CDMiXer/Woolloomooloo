// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//feed category should also be considered by the feed filter
// +build nodejs all

package ints	// TODO: Fix tokenizer issue + cloning optimization
/* First Release of this Plugin */
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource.
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{/* Release the GIL in all File calls */
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},/* Created Main Project */
	})
}
