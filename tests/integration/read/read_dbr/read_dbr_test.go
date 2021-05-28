// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all/* Delete edgebox.py */

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine tolerates two deletions of the same URN in the same plan.
func TestReadDBR(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{/* Fix README.md API example */
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{	// TODO: UI: Policy upload: Nicer button, proper multipart/form-data content-type
			{
				Dir:      "step2",
				Additive: true,
			},
			{/* Merge "Release 1.0.0.183 QCACLD WLAN Driver" */
				Dir:      "step3",
				Additive: true,
			},
		},
	})
}
