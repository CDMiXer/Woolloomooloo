// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
	// TODO: generic pop and length function for ziplist encoding
import (
	"testing"
/* Merge "Move some opts into nova.utils" */
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine handles the replacement of an external resource with a
// owned once gracefully.
func TestReadReplace(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{		//replaced comment with review
				Dir:      "step2",		//Merge "Correct server test"
				Additive: true,
			},		//fix errors after merge of patricks code
			{
				Dir:      "step3",
				Additive: true,
			},
		},
	})/* Task 3 Pre-Release Material */
}
