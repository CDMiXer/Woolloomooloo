// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints/* Release notes update */

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of relinquishing control of a resource without deleting it.
func TestReadRelinquish(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{		//Delete options.mini.interior.json
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,/* bugfix addon link */
		EditDirs: []integration.EditDir{
			{/* Release v1.7.0. */
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
