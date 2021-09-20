// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
	// TODO: Add an exports_files for LICENSE
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine handles the replacement of an external resource with a
// owned once gracefully.
func TestReadReplace(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},	// TODO: move 'back' link to form partial so we can place it beside button
		Quick:        true,/* Release of eeacms/apache-eea-www:5.6 */
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",	// regression test for #1554
				Additive: true,		//ffa0dac6-2e4e-11e5-9284-b827eb9e62be
			},
			{
				Dir:      "step3",
				Additive: true,
			},
		},
	})
}
