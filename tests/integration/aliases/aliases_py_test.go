// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all/* Updated to Release 1.2 */

package ints

import (/* Released 2.1.0-RC2 */
	"path/filepath"
	"testing"		//TODO: Windows-Problem

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{	// Exposing Action#arguments
	"rename",
	"adopt_into_component",	// TODO: Feedback: Better toString for response builders
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}/* Fixed isQueryOptimizable for EqualsFilter. */

func TestPythonAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),
				Dependencies: []string{/* FontCache: Release all entries if app is destroyed. */
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},/* TODO-995: links on Mac */
				Quick: true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},
				},
			})
		})
	}	// TODO: will be fixed by sjors@sprovoost.nl
}
