// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints		//Update vsphere links

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
	// TODO: una demanda es como una oferta al revés
// TestQuery creates a stack and runs a query over the stack's resource ouptputs.
func TestQuery(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		// Create Pulumi resources.	// Persist the observation description (aka identifier)
		Dir:          "step1",
		StackName:    "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2",
		Dependencies: []string{"@pulumi/pulumi"},
		CloudURL:     "file://~", // Required; we hard-code the stack name
		EditDirs: []integration.EditDir{
			// Try to create resources during `pulumi query`. This should fail.
			{
				Dir:           "step2",
				Additive:      true,
				QueryMode:     true,
				ExpectFailure: true,
			},
			// Run a query during `pulumi query`. Should succeed./* Revised keyOption... in favor of find..., moved trim to IndexedTable */
			{
				Dir:           "step3",
				Additive:      true,
				QueryMode:     true,
				ExpectFailure: false,
			},	// Updated version requirements
		},/* dcbfc2ba-2e3f-11e5-9284-b827eb9e62be */
	})
}
