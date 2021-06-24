// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build go all

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{	// TODO: will be fixed by remco@dutchcoders.io
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",	// state: fix agent version error messages
}

func TestGoAliases(t *testing.T) {
	for _, dir := range dirs {	// TODO: will be fixed by xaber.twt@gmail.com
		d := filepath.Join("go", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),
				Dependencies: []string{
					"github.com/pulumi/pulumi/sdk/v2",	// TODO: will be fixed by steven@stebalien.com
				},
				Quick: true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						ExpectNoChanges: true,		//Test Address Prefix changed to T
						Additive:        true,
					},
				},
			})
		})
	}
}
