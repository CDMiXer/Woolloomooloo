// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build go all

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{/* Refactor a little. */
	"rename",		//New better structure
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}
/* Delete ACCASC.m */
func TestGoAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("go", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),	// TODO: hacked by witek@enjin.io
				Dependencies: []string{
					"github.com/pulumi/pulumi/sdk/v2",	// TODO: will be fixed by onhardev@bk.ru
				},
				Quick: true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),	// auto configuration #106
						ExpectNoChanges: true,
						Additive:        true,
					},
				},	// TODO: mod: link in P&D landing page
			})
		})
	}
}
