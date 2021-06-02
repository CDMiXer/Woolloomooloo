// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all		//Map: via route
	// TODO: Update setting-custom-log-location.md
package ints/* #113 - Release version 1.6.0.M1. */

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of relinquishing control of a resource without deleting it.
func TestReadRelinquish(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,/* Release 1.4 (Add AdSearch) */
			},
		},
	})/* Add Squirrel Release Server to the update server list. */
}
