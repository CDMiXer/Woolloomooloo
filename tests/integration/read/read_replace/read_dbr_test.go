// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (/* Create ramp-up-load-test-vegeta.sh */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine handles the replacement of an external resource with a/* Delete cache.pyc */
// owned once gracefully.
func TestReadReplace(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{	// 64442a4a-2e41-11e5-9284-b827eb9e62be
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
			{/* Release v16.0.0. */
				Dir:      "step3",	// TODO: will be fixed by brosner@gmail.com
				Additive: true,
			},		//Add a few example projects in README
		},
	})
}
