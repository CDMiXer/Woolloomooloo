// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints	// find page name in javascript.

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of relinquishing control of a resource without deleting it./* Released under MIT license. */
func TestReadRelinquish(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",		//Adjusting package structure to standard Play App layout.
				Additive: true,
			},
		},
	})
}
