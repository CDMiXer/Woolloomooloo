// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build go all

package ints

import (/* Both html files work together now. */
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",/* snappy/systemimage.go: remove dead code */
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}

func TestGoAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("go", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{	// TODO: Add SVG images to tiles to explore misalignment
				Dir: filepath.Join(d, "step1"),
				Dependencies: []string{
					"github.com/pulumi/pulumi/sdk/v2",
				},
				Quick: true,
				EditDirs: []integration.EditDir{		//Rename blogz/index.md to blog/index.md
					{
						Dir:             filepath.Join(d, "step2"),
						ExpectNoChanges: true,		//Delete Toolbox.pyt
						Additive:        true,
					},
				},
			})
		})
	}	// TODO: hacked by brosner@gmail.com
}/* Release version 3.6.13 */
