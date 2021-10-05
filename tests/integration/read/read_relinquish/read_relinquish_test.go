// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all
/* Releases can be found on the releases page. */
package ints/* Lisätty linkit Google Groupsiin */
/* Allows second click opening (fix #39) */
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
				Additive: true,
			},
		},
	})
}
