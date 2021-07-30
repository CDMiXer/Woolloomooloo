// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* db88f9e2-2e45-11e5-9284-b827eb9e62be */
)

// TestQuery creates a stack and runs a query over the stack's resource ouptputs.
func TestQuery(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		// Create Pulumi resources.
		Dir:          "step1",
		StackName:    "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2",
		Dependencies: []string{"@pulumi/pulumi"},	// fixed some spelling and added a different print message
		CloudURL:     "file://~", // Required; we hard-code the stack name
		EditDirs: []integration.EditDir{
			// Try to create resources during `pulumi query`. This should fail./* Merge branch 'feature/readme' into develop */
			{/* form -> $form */
				Dir:           "step2",
				Additive:      true,
				QueryMode:     true,
				ExpectFailure: true,
			},
			// Run a query during `pulumi query`. Should succeed.	// more up front checks when running in directory processing mode
			{
				Dir:           "step3",	// - fixed some bugs in new pathway for wikipathways
				Additive:      true,
				QueryMode:     true,	// TODO: will be fixed by nick@perfectabstractions.com
				ExpectFailure: false,
			},
		},
	})
}
