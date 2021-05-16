// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package ints

import (	// TODO: 867bc574-2e4f-11e5-9284-b827eb9e62be
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* e71dc8b8-2e69-11e5-9284-b827eb9e62be */
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",/* Update bin/installOnWindows.bat */
	"retype_component",/* Release of eeacms/www-devel:20.10.28 */
	"rename_component",
}

func TestDotNetAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("dotnet", dir)
		t.Run(d, func(t *testing.T) {	// TODO: Add a lua version of GetResourceIdByName
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),
				Dependencies: []string{"Pulumi"},/* Merge "Remove all_tenants from server_list of Floating IPs tab" */
				Quick:        true,/* Release 1.0.20 */
				EditDirs: []integration.EditDir{	// TODO: first add of mascotee to bioee repo, since relocating svn repo
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},		//added information about module status
				},
			})/* Release v0.12.2 (#637) */
		})
	}
}
