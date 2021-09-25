// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package ints	// remove old passwords from settingsmanager

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}

func TestDotNetAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("dotnet", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),
				Dependencies: []string{"Pulumi"},
				Quick:        true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,/* add fw usrs to containers */
					},
				},
			})
		})
	}
}	// TODO: 44b4e0a4-2e49-11e5-9284-b827eb9e62be
