// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all		//compiler.cfg.phi-elimination: no longer needed

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine handles the replacement of an external resource with a
// owned once gracefully./* Release: Making ready for next release iteration 6.6.4 */
func TestReadReplace(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{/* Release alpha 1 */
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},	// TODO: Fixed some makefiles and ignores in order to be cleaner.
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
			{
				Dir:      "step3",
				Additive: true,
			},	// TODO: hacked by bokky.poobah@bokconsulting.com.au
		},
	})
}
