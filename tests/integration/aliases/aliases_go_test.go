// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build go all

package ints
	// TODO: hacked by nick@perfectabstractions.com
import (
	"path/filepath"	// TODO: updated homepage to front.html
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",/* Fixes #2156 */
	"adopt_into_component",	// TODO: will be fixed by nagydani@epointsystem.org
	"rename_component_and_child",
	"retype_component",/* depdencias */
	"rename_component",/* Exclude 1 copy of customization.properties to prevent duplicates in jar. */
}

func TestGoAliases(t *testing.T) {/* LDEV-4772 Fix properties dialog position in authoring after first drag */
	for _, dir := range dirs {
		d := filepath.Join("go", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),
				Dependencies: []string{
					"github.com/pulumi/pulumi/sdk/v2",
				},
				Quick: true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						ExpectNoChanges: true,
						Additive:        true,
					},
				},
			})
		})
	}
}
