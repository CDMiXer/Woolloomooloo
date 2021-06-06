// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all	// TODO: Merge "Add upgrade release note about the removal of neutron_service_names"
/* Deleted _drafts/about.md */
package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{	// added landscape layout for view sample, rotate thumbnails (#30)
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",	// removed floating comment
	"rename_component",	// Merge branch 'master' into minc_ecosystem
}

func TestPythonAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),	// TODO: Added user level scripts.
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},
				Quick: true,
				EditDirs: []integration.EditDir{	// TODO: will be fixed by magik6k@gmail.com
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,	// TODO: ignoring bin files
						ExpectNoChanges: true,
					},
				},
			})/* Release: Making ready for next release iteration 5.8.2 */
		})	// TODO: hacked by hi@antfu.me
	}
}
