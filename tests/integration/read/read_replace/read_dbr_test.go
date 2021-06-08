// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"testing"	// TODO: Edited about page

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
	// importing everything important
// Test that the engine handles the replacement of an external resource with a
// owned once gracefully.
func TestReadReplace(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{/* general code cleanup and integration of Titan */
				Dir:      "step2",
				Additive: true,
			},
			{	// depuracion de filtros en detalle vacunacion
				Dir:      "step3",
				Additive: true,
			},
		},
	})
}
