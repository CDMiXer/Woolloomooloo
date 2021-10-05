// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build go all
	// TODO: fixed compiled PMAP problem by narrowing the scope of the issue.
package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Release version 1.6.0.M1 */
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",/* Release of eeacms/www-devel:18.3.14 */
	"retype_component",
	"rename_component",
}

func TestGoAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("go", dir)/* Release Candidate 4 */
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
